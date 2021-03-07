// Package main allows generating requests to test api
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const server = "http://localhost:8080"

func main() {
	var maxSize, n int
	var concurrent bool
	flag.IntVar(&n, "n", 10, "Number of POST requests to generate")
	flag.IntVar(&maxSize, "maxsize", 10, "Number of POST requests to generate")
	flag.BoolVar(&concurrent, "concurrent", true, "Run concurrent requests")
	flag.Parse()
	if maxSize < 4 {
		maxSize = 4
	}
	if n < 0 {
		n = 0
	}

	// Random seed to generate different values on every execution
	rand.Seed(time.Now().UnixNano())
	// rand.Seed(1) // to test speed difference of postMutants and postMutantsConcurrent

	// Get ammount of runs from 1st command line argument, default 10 if no argument supplied

	getStats()
	start := time.Now()
	if concurrent {
		postMutantsConcurrent(maxSize, n)
	} else {
		postMutants(maxSize, n)
	}
	elapsed := time.Since(start)
	getStats()
	fmt.Printf("postMutant(%d) took %s\n", n, elapsed)
}

func buildDna(maxSize int) []string {
	// Valid gene pool
	genes := []string{"A", "T", "C", "G"}
	size := rand.Intn(maxSize)
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
		fmt.Printf("HTTP request failed. Error: %s\n", err)
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
		fmt.Printf("HTTP request failed. Error: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Stats:", string(data))
	}
}

func postMutantsConcurrent(maxSize int, n int) {
	sem := make(chan struct{}, 100)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		select {
		case sem <- struct{}{}:
			go func() {
				postMutant(buildDna(maxSize))
				<-sem
				wg.Done()
			}()
		default:
			postMutant(buildDna(maxSize))
			wg.Done()
		}
	}
	wg.Wait()
}

func postMutants(maxSize int, n int) {
	for i := 0; i < n; i++ {
		postMutant(buildDna(maxSize))
	}
}
