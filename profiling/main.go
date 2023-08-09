package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import for pprof profiling endpoints
	"time"
)

func compute(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	return result
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Profiling server started. Access the web UI at http://localhost:6060/debug/pprof/")
	time.Sleep(1 * time.Hour) // Keep the program running for a while
}
