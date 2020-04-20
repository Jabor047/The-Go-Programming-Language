package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenServe("localhost:8000", nil))
}

// handler echoes the path component of the requested url
func handler(w http.ResponseWriter, r *http.Request){
	mu.lock()
	count++
	mu.unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request){
	mu.lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}