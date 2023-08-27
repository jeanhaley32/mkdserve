package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/sync/semaphore"
)

func ServePages(sem *semaphore.Weighted) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if !sem.TryAcquire(1) {
			log.Printf("Too many connections from %s ", r.RemoteAddr)
			http.Error(w, "Too many connections", http.StatusTooManyRequests)
			return
		}
		defer sem.Release(1)
		log.Printf("Serving request from %s %s\n", r.RemoteAddr, r.URL.Path)
		HTMLHandler(w, r)
	})
}

func HTMLHandler(w http.ResponseWriter, r *http.Request) {

	content, err := os.ReadFile(page)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		http.Error(w, "File not found.", 404)
		return
	}

	// Set the Content-Type to HTML and write the response to the client.
	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
}
