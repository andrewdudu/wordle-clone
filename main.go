package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

var sliceWords []string
var mapWords words
var todayAnswer string

func main() {
	todayAnswer = "broke"
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		sliceWords = append(sliceWords, scanner.Text())
	}

	mapWords = sliceToMap(sliceWords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/guess", GuessHandler)
	http.ListenAndServe(":8080", nil)
}
