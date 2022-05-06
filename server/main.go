package main

import (
	"errors"
	"file_transfer/core"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"syscall"
	"time"
)

const BUFFERSIZE = 1024

var filePath = make(chan string, 10)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go uploadServer(&wg)
	go downloadServer(&wg)
	wg.Wait()

	fmt.Println("server finished!")
}

func uploadServer(wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		port        = 8877
		http2       = false
		key         = ""
		certificate = ""
		server      core.Server
		err         error
	)

	switch {
	case http2:
		if key == "" || certificate == "" {
			must(errors.New(
				"http2 requires key and certificate to be specified"))
		}

		http2Server, err := core.NewServerH2(core.ServerH2Config{
			Port:        port,
			Certificate: certificate,
			Key:         key,
		})
		must(err)
		server = &http2Server
	default:
		grpcServer, err := core.NewServerGRPC(core.ServerGRPCConfig{
			Port:        port,
			Certificate: certificate,
			Key:         key,
			FilePath:    filePath,
		})
		must(err)
		server = &grpcServer
	}

	err = server.Listen()
	must(err)
	defer server.Close()

}

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case filepath := <-filePath:
			for cli := range clients {
				cli <- filepath
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			log.Println("one connection leave.")
			delete(clients, cli)
			close(cli)
		}
	}
}

func downloadServer(wg *sync.WaitGroup) {
	defer wg.Done()
	server, err := net.Listen("tcp", "localhost:27001")
	if err != nil {
		fmt.Println("Error listetning: ", err)
		os.Exit(1)
	}
	go broadcaster()
	defer server.Close()
	fmt.Println("Server started! Waiting for connections...")
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatalf("Error: %v", err)
			os.Exit(1)
		}
		log.Println("Client connected")
		go handleConn(connection)

	}
}

func handleConn(conn net.Conn) {
	defer trace("handleConn:" + conn.RemoteAddr().String())()

	ch := make(chan string)

	go sendFileToClient(conn, ch)
	entering <- ch

	//
	one := make([]byte, 1)
	for {
		conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		_, err := conn.Read(one)
		if err == io.EOF {
			log.Println("linux connect close ")
			goto LEAVING
		} else if opErr, ok := err.(*net.OpError); ok {
			if syscallErr, ok := opErr.Err.(*os.SyscallError); ok {
				if se, ok := syscallErr.Err.(syscall.Errno); ok && se == syscall.WSAECONNRESET {
					log.Printf("windows connect close %s", err)
					goto LEAVING
				}
			}
		}
	}

LEAVING:
	leaving <- ch
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func sendFileToClient(connection net.Conn, ch chan string) {
	log.Println("A client has connected!")

	for filepath := range ch {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}
		fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
		fileName := fillString(fileInfo.Name(), 64)
		fmt.Printf("Sending filename and filesize!\nfilesize:%s,filename:%s\n", []byte(fileSize), []byte(fileName))
		connection.Write([]byte(fileSize))
		connection.Write([]byte(fileName))
		sendBuffer := make([]byte, BUFFERSIZE)
		fmt.Println("Start sending file! ")
		for {
			_, err = file.Read(sendBuffer)
			if err == io.EOF {
				break
			}
			_, err = connection.Write(sendBuffer)
			if err == io.EOF {

			}
		}
		file.Close()
	}
	//fmt.Println("File has been sent, closing connection!")
	return
}

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

func must(err error) {
	if err == nil {
		return
	}

	log.Printf("ERROR: %+v\n", err)
	os.Exit(1)
}
