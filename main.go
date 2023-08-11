package main

import (
	"flag"
	"log"
	"net/http"

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
	sem := semaphore.NewWeighted(ConnectionLimit)

	ServeImages(sem)
	ServePages(sem)

	log.Printf("Starting server on http://%s\n", socket)
	if err := http.ListenAndServe(socket, nil); err != nil {
		log.Fatal(err)
	}
}
