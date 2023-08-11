package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/semaphore"
)

func ServeImages(sem *semaphore.Weighted) {
	http.HandleFunc("/image/", func(w http.ResponseWriter, r *http.Request) {
		// Acquire a semaphore slot
		if !sem.TryAcquire(1) {
			log.Printf("Too many connections from %s ", r.RemoteAddr)
			http.Error(w, "Too many connections", http.StatusTooManyRequests)
			return
		}
		// Release the semaphore slot when we're done
		defer sem.Release(1)
		log.Printf("Serving request from %s %s\n", r.RemoteAddr, r.URL.Path)
		imageName := strings.TrimPrefix(r.URL.Path, "/image/")
		imgPath := filepath.Join("image", imageName)
		if imageName == "" {
			GetImageLinkPage(ContentScraper([]string{".jpeg", ".png", ".gif"}, "image"))
		}

		// Open the image file
		img, err := os.Open(imgPath)
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}
		defer img.Close()

		// Set the content type header based on the file extension
		contentType := "image/jpeg" // Default to JPEG
		if strings.HasSuffix(imageName, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(imageName, ".gif") {
			contentType = "image/gif"
		}
		w.Header().Set("Content-Type", contentType)

		// Copy the image file to the response writer
		if _, err := io.Copy(w, img); err != nil {
			log.Printf("failed to copy image, %s", err.Error())
			http.Error(w, "Internal server error", 500)
		}
	})
}

type Link struct {
	URL   string
	Label string
}

func GetImageLinkPage(links []string) []byte {
	var b bytes.Buffer
	// Define the template
	tmpl := template.Must(template.New("links").Parse(`
        <!DOCTYPE html>
        <html>
            <head>
                <title>Images</title>
            </head>
            <body>
                <h1>Image Links</h1>
                <ul>
                    {{range .}}
                        <li><a href="{{.URL}}">{{.Label}}</a></li>
                    {{end}}
                </ul>
            </body>
        </html>
    `))

	// Define the data for the template
	var data []Link
	for _, link := range links {
		data = append(data, Link{link, filepath.Base(link)})
	}

	err := tmpl.Execute(&b, data)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}
