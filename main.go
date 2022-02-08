package main

import (
	"encoding/json"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome")
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8080", nil)
}
