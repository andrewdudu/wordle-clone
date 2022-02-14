package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"fmt"
)

var sliceWords []string
var mapWords words
var todayAnswer string

func main() {
	todayAnswer = "broke"
	readWords()
	
	http.HandleFunc("/guess", GuessHandler)
	http.HandleFunc("/shuffle", ShuffleHandler)
	http.ListenAndServe(":8080", nil)
}

func readWords() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		sliceWords = append(sliceWords, scanner.Text())
		fmt.Println(scanner.Text())
	}

	mapWords = sliceToMap(sliceWords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
