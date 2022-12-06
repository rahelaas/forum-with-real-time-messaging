package main

import "net/http"

func ConfigHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	return w
}
