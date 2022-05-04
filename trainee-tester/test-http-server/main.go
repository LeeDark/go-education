package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/hello", hello)

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println(runtime.NumCPU())
	log.Println("Listening...")
	server.ListenAndServe()
}

// Active endpoint = HTTP Server gives HTML page
func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `
	<HTML>
	<HEAD><TITLE>HTTP Server</TITLE></HEAD>
	<BODY><P align='center'>Hello! Welcome to HTTP Server written on Golang!</P></BODY>
	</HTML>
	`)
}

// Passive endpoint = HTTP Server gives JSON (XML) data = Frontend uses this JSON (XML) data
func ping(w http.ResponseWriter, req *http.Request) {
	//log.Println("Got ping")
	fmt.Fprintf(w, "{ \"answer\": \"pong\" }")
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)

	name := req.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "{ \"answer\": \"error\" }")
	}

	fmt.Fprintf(w, "{ \"answer\": \"%s\" }", name)
}
