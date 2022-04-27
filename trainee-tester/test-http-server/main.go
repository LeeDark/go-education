package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/ping", ping)

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

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
	fmt.Fprintf(w, "{ \"answer\": \"pong\"}")
}
