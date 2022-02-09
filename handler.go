package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type GuessRequest struct {
	Request string `json:"request"`
}

type GuessResponse struct {
	Result string `json:"result"`
	Valid  bool   `json:"isValid"`
}

func ConstructResponse() GuessResponse {
	res := GuessResponse{
		Valid: true,
	}

	return res
}

func GuessHandler(w http.ResponseWriter, r *http.Request) {
	var guessRequest GuessRequest
	res := ConstructResponse()
	result := "-----"
	ConstructRequest(&guessRequest, w, r)

	word := guessRequest.Request
	_, isFound := mapWords[word]
	if len(word) != 5 {
		res.Valid = false
	}

	if isFound {
		result = CalculateAnswer(word, result)
	}

	fmt.Println(isFound)
	fmt.Println(word)

	res.Result = result

	fmt.Println(guessRequest)
	json.NewEncoder(w).Encode(res)
}

func CalculateAnswer(ans string, result string) string {
	for i, char := range ans {
		if ans[i] == todayAnswer[i] {
			result = replaceAtIndex(result, '1', i)
		} else if strings.Contains(todayAnswer, string(char)) {
			result = replaceAtIndex(result, '0', i)
		}
	}

	return result
}
