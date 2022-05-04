package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	router "github.com/gorilla/mux"
)

func main() {
	mux := router.NewRouter()

	mux.HandleFunc("/", home)

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
