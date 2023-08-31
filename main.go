package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	ip, socket, page, key, csr string
	ConnectionLimit            int64
)

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "IP address to listen on, defaults to 127.0.0.1 if not set")
	flag.StringVar(&page, "main", "main.html", "Main markdown file to serve:, defaults to main.html if not set")
	flag.StringVar(&csr, "csr", "", "SSL certificate file, defaults to blank if not set")
	flag.StringVar(&key, "key", "", "SSL key file, defaults to blank if not set")
	flag.Parse()
}

func main() {
	// Handle main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If root path, serve main page
		if r.URL.Path == "/" {
			http.ServeFile(w, r, page)
		}
		// if not root path, serve requested page
		target := r.URL.Path[0:]
		http.ServeFile(w, r, "./pages/"+target+".html")
	})
	// Handle Image subdirectory
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	// Handle Assets subdirectory
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Printf("Starting server on http://%s\n", socket)

	// If SSL certificate and key are provided, start server with TLS
	if csr != "" && key != "" {
		socket := ip + ":443"
		log.Printf("Starting Encrtyped TLS server on https://%s\n", socket)
		if err := http.ListenAndServeTLS(socket, csr, key, nil); err != nil {
			log.Fatal(err)
		}
	} else {
		socket := ip + ":80"
		// If SSL certificate and key are not provided, start server without TLS
		log.Printf("Starting http server on http://%s\n", socket)
		if err := http.ListenAndServe(socket, nil); err != nil {
			log.Fatal(err)
		}
	}
}
