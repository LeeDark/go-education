package main

import (
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	mux := http.NewServeMux()
	routes(mux)

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
