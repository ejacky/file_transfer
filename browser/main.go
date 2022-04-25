// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"errors"
	"file_transfer/core"
	"fmt"
	"golang.org/x/net/context"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadFile)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)

}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	var (
		chunkSize       = 1 << 12
		http2           = false
		address         = "localhost:8877"
		rootCertificate = ""
		compress        = false
		client          core.Client
	)

	if address == "" {
		must(errors.New("address"))
	}

	switch {
	case http2:
		if rootCertificate == "" {
			must(errors.New("http2 requires root-certificate to be supplied"))
		}

		if !strings.HasPrefix(address, "https://") {
			address = "https://" + address
		}

		http2Client, err := core.NewClientH2(core.ClientH2Config{
			Address:         address,
			RootCertificate: rootCertificate,
		})
		must(err)
		client = &http2Client
	default:
		grpcClient, err := core.NewClientGRPC(core.ClientGRPCConfig{
			Address:         address,
			RootCertificate: rootCertificate,
			Compress:        compress,
			ChunkSize:       chunkSize,
		})
		must(err)
		client = &grpcClient
	}

	stat, err := client.UploadFile(context.Background(), r)
	must(err)
	defer client.Close()

	fmt.Printf("%d\n", stat.FinishedAt.Sub(stat.StartedAt).Nanoseconds())

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

//!-handler

func must(err error) {
	if err == nil {
		return
	}

	fmt.Printf("ERROR: %+v\n", err)
	os.Exit(1)
}
