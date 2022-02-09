package main

import (
	"encoding/json"
	"net/http"
)

type request interface {
	GuessRequest
}

func ConstructRequest[R request](req *R, w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}
