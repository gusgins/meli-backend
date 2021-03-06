// Package main allows generating requests to test api
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const server = "http://localhost:8080"

func main() {
	// valid gene pool
	genes := []string{"A", "T", "C", "G"}

	// Random seed to generate different values on every execution
	rand.Seed(time.Now().UnixNano())
	N := 10
	// Get ammount of runs from 1st command line argument, default 10 if no argument supplied
	if len(os.Args) > 1 {
		var err error
		N, err = strconv.Atoi(os.Args[1])
		// If any error, default to 10 runs
		if err != nil {
			N = 10
		}
	}

	for i := 0; i < N; i++ {
		dna := buildDna(genes)
		postMutant(dna)
	}
	getStats()
}

func buildDna(genes []string) []string {
	size := rand.Intn(15) + 4
	dna := make([]string, size)
	for j := 0; j < size; j++ {
		str := ""
		for k := 0; k < size; k++ {
			str += genes[rand.Intn(4)]
		}
		dna[j] = str
	}
	return dna
}

func postMutant(dna []string) {
	json, _ := json.Marshal(map[string][]string{"dna": dna})
	response, err := http.Post(server+"/mutant", "application/json", bytes.NewBuffer(json))
	if err != nil {
		fmt.Printf("HTTP request failed. Error: %s\n", err.Error())
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		_ = data
		// Uncomment to print the request body and response
		//fmt.Println(string(json), string(data))
	}
}

func getStats() {
	response, err := http.Get(server + "/stats")
	if err != nil {
		fmt.Printf("HTTP request failed. Error: %s\n", err.Error())
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Stats:", string(data))
	}
}
