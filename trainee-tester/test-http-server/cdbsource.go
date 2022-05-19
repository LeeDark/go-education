package main

import "net/http"

func cdbSourceHandler(w http.ResponseWriter, req *http.Request) {
	sendJSON(http.StatusOK, w, map[string]string{"answer": "pong"})
}
