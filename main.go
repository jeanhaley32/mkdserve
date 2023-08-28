package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port, ip, socket, page string
	ConnectionLimit        int64
)

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "IP address to listen on, defaults to 127.0.0.1 if not set")
	flag.StringVar(&port, "port", "8080", "Port to listen on, defaults to 8080 if not set")
	flag.StringVar(&page, "main", "main.html", "Main markdown file to serve:, defaults to main.html if not set")
	flag.Int64Var(&ConnectionLimit, "limit", 100, "Connection limit, defaults to 100 if not set")
	flag.Parse()
	socket = ip + ":" + port
}

func main() {

	// Handle main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, page)
		}
		target := r.URL.Path[1:]
		http.ServeFile(w, r, "/pages/"+target+".html")
	})
	// Handle Image subdirectory
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	// Handle Assets subdirectory
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Printf("Starting server on http://%s\n", socket)
	if err := http.ListenAndServe(socket, nil); err != nil {
		log.Fatal(err)
	}
}
