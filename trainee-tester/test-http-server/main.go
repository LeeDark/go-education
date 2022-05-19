package main

import (
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	// router := SetRoutes()
	mux := setEndpoints()

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Num CPU:", runtime.NumCPU())
	log.Println("Listening...")
	server.ListenAndServe()
}
