// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"file_transfer/core"
	"file_transfer/messaging"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	address         string
	serverAddress   string
	chunkSize       = 1 << 12
	http2           = false
	rootCertificate = ""
	compress        = false
	client          core.Client
)

func init() {
	cfg, err := ini.Load("conf/app.conf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	address = cfg.Section("browser").Key("address").String()
	serverAddress = cfg.Section("server").Key("address").String()

}

func main() {

	// 页面
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		file, err := os.Open("index.html")
		if err != nil {
			log.Fatal(err)
			return
		}
		content, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = writer.Write(content)
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	// 接口
	http.HandleFunc("/file/mpupload/init", mpuploadInit)
	http.HandleFunc("/file/mpupload/uppart", mpuploadUppart)
	http.HandleFunc("/file/mpupload/complete", mpuploadComplete)
	//http.HandleFunc("/file/mpupload/cancel", uploadFile)

	log.Fatal(http.ListenAndServe(address, nil))
}

func mpuploadInit(w http.ResponseWriter, r *http.Request) {

	fileHash := r.FormValue("filehash")
	fileSize, err := strconv.Atoi(r.FormValue("filesize"))
	if err != nil {
		log.Fatal(err)
	}

	grpcClient, err := core.NewClientGRPC(core.ClientGRPCConfig{
		Address:         serverAddress,
		RootCertificate: rootCertificate,
		Compress:        compress,
		ChunkSize:       chunkSize,
	})
	must(err)
	client = &grpcClient

	data, err := grpcClient.Client.Init(context.Background(), &messaging.InitReq{
		FileHash: fileHash,
		FileSize: uint64(fileSize),
	})
	if err != nil {
		log.Fatal(err)
	}

	resultOmitEmptyJson(w, data)
}

func mpuploadUppart(w http.ResponseWriter, r *http.Request) {

	stat, err := client.UploadFile(context.Background(), r)
	must(err)

	fmt.Printf("%d\n", stat.FinishedAt.Sub(stat.StartedAt).Nanoseconds())

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded Chunk\n")
}

func mpuploadComplete(w http.ResponseWriter, r *http.Request) {
	defer client.Close()
	fileSize, err := strconv.Atoi(r.FormValue("filesize"))
	if err != nil {
		log.Fatal(err)
	}
	grpcClient, ok := client.(*core.ClientGRPC)
	if ok {
		data, err := grpcClient.Client.Complete(context.Background(), &messaging.CompleteReq{
			//todo 校验必填
			FileHash: r.FormValue("filehash"),
			FileName: r.FormValue("filename"),
			UploadID: r.FormValue("uploadid"),
			FileSize: uint64(fileSize),
		})
		if err != nil {
			log.Fatal(err)
		}

		resultOmitEmptyJson(w, data)
	} else {
		fmt.Fprintf(w, "doesn't support protocal type \n")
	}
}

func resultOmitEmptyJson(w http.ResponseWriter, data proto.Message) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	m := jsonpb.Marshaler{EmitDefaults: true}
	err := m.Marshal(w, data.(proto.Message))
	if err != nil {
		log.Fatal("序列化 proto 失败:", err)
	}
}

func must(err error) {
	if err == nil {
		return
	}

	fmt.Printf("ERROR: %+v\n", err)
	os.Exit(1)
}
