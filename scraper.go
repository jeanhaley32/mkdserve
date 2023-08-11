// Scraper is used to scrap folders for servable content
package main

import (
	"os"
	"path/filepath"
)

// Content Scraper takes in a list of extensions and a path to walk.
// It looks for any files with these extension, and return them as a list of strings.
func ContentScraper(e []string, l string) []string {
	var content []string
	// Walk the path and append the files with the extension to the content list
	err := filepath.Walk(l, func(path string, info os.FileInfo, err error) error {
		for _, ext := range e {
			if filepath.Ext(path) == ext {
				content = append(content, path)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return content
}
