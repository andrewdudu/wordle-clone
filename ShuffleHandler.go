package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
)

type ShuffleRequest struct {
	Key string `json:"key"`
}

type ShuffleResponse struct {
	Success bool `json:"success"`
}

func ConstructShuffleResponse() ShuffleResponse {
	res := ShuffleResponse{
		Success: true,
	}

	return res
}

func ShuffleHandler(w http.ResponseWriter, r *http.Request) {
	var shuffleRequest ShuffleRequest
	res := ConstructShuffleResponse()
	ConstructRequest(&shuffleRequest, w, r)
	ShuffleAnswer()

	json.NewEncoder(w).Encode(res)
}

func ShuffleAnswer() {
	todayAnswer = sliceWords[rand.Intn(len(sliceWords))]
}