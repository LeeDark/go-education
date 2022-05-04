package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func setEndpoints() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/hello", hello)
	return mux
}

func sendJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// TODO: add error handling
	return json.NewEncoder(w).Encode(&data)
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
	sendJSON(w, struct {
		Answer string `json:"answer"`
	}{Answer: "pong"})
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)

	name := req.URL.Query().Get("name")
	if name == "" {
		sendJSON(w, struct {
			Answer string `json:"answer"`
		}{Answer: "Param name was not found"})
		return
	}

	sendJSON(w, struct {
		Answer string `json:"answer"`
	}{Answer: name})
}
