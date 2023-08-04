package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
	"golang.org/x/sync/semaphore"
)

var (
	port, ip, socket, page string
	ConnectionLimit        int64
)

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "IP address to listen on, defaults to 127.0.0.1 if not set")
	flag.StringVar(&port, "port", "8080", "Port to listen on, defaults to 8080 if not set")
	flag.StringVar(&page, "main", "main.md", "Main markdown file to serve:, defaults to main.md if not set")
	flag.Int64Var(&ConnectionLimit, "limit", 100, "Connection limit, defaults to 100 if not set")
	flag.Parse()
	socket = ip + ":" + port
}

func main() {
	fmt.Println(ConnectionLimit)
	sem := semaphore.NewWeighted(ConnectionLimit)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !sem.TryAcquire(1) {
			log.Printf("Too many connections from %s ", r.RemoteAddr)
			http.Error(w, "Too many connections", http.StatusTooManyRequests)
			return
		}
		defer sem.Release(1)
		log.Printf("Serving request from %s: %s\n", r.RemoteAddr, r.URL.User)
		MarkdownHandler(w, r)
	})

	log.Printf("Starting server on http://%s\n", socket)
	if err := http.ListenAndServe(socket, nil); err != nil {
		log.Fatal(err)
	}
}

func MarkdownHandler(w http.ResponseWriter, r *http.Request) {

	content, err := os.ReadFile(page)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		http.Error(w, "File not found.", 404)
		return
	}

	output := blackfriday.Run(content)

	// Set the Content-Type to HTML and write the response to the client.
	w.Header().Set("Content-Type", "text/html")
	w.Write(output)
	log.Printf("success")
}
