package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/LeeDark/go-education/trainee-tester/test-http-server/middleware"
)

func setEndpoints() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/hello", middleware.AuthName(hello))
	mux.HandleFunc("/timeout", timeout)

	mux.HandleFunc("/cdbsource", cdbSourceHandler)

	handler := middleware.Logging(mux)

	return handler
}

func sendJSON(statusCode int, w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// TODO: add error handling
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		log.Println("ERROR:", err)
	}
	return err
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
	if req.Method != "GET" {
		sendJSON(http.StatusMethodNotAllowed, w, map[string]string{"answer": "Method should be GET"})
		return
	}
	sendJSON(http.StatusOK, w, map[string]string{"answer": "pong"})
}

func hello(w http.ResponseWriter, req *http.Request) {
	// log.Println(req.URL)
	name := req.URL.Query().Get("name")
	if name == "" {
		// sendJSON(w, struct {
		// 	Answer string `json:"answer"`
		// }{Answer: "Param name was not found"})
		sendJSON(http.StatusBadRequest, w, map[string]string{"answer": "Param name was not found"})
		return
	}

	// sendJSON(w, struct {
	// 	Answer string `json:"answer"`
	// }{Answer: name})
	sendJSON(http.StatusOK, w, map[string]string{"answer": name})
}

func randomTimeout(from, to int) time.Duration {
	return time.Duration(rand.Intn(to-from)+from) * time.Millisecond
}

func timeout(w http.ResponseWriter, req *http.Request) {
	// randomizer 0.5-1 sec
	time.Sleep(randomTimeout(500, 1000))

	sendJSON(http.StatusOK, w, map[string]string{"answer": "pong"})
}
