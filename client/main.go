package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const BUFFERSIZE = 1024

func main() {
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	log.Println("Connected to server, start receiving the file name and file size")
	for {
		bufferFileName := make([]byte, 64)
		bufferFileSize := make([]byte, 10)

		connection.Read(bufferFileSize)
		fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

		connection.Read(bufferFileName)
		fileName := strings.Trim(string(bufferFileName), ":")

		log.Printf("download filename and filesize!\nfilesize:%s,filename:%s\n", bufferFileSize, bufferFileName)

		newFile, err := os.Create(fileName)

		if err != nil {
			panic(err)
		}

		var receivedBytes int64

		for {
			if (fileSize - receivedBytes) < BUFFERSIZE {
				io.CopyN(newFile, connection, (fileSize - receivedBytes))
				connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
				break
			}
			io.CopyN(newFile, connection, BUFFERSIZE)
			receivedBytes += BUFFERSIZE
		}

		newFile.Close()
	}

	log.Println("Received file completely!")
}
