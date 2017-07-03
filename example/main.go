package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/adrianplavka/bulk"
)

/*
	Author: 	Adrian Plavka
	Contact: 	adrian.plavka@gmail.com
	Reference: 	https://www.github.com/adrianplavka/bulk
*/

func main() {
	// Bulker is a type, that is defined by http.Client.
	bulker := bulk.DefaultBulker

	// Open a CSV file for read-only.
	path, _ := filepath.Abs("example/urls.csv")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("failed while opening a file: ", err)
	}

	// Create a LineDecoder with a Status channel.
	// File is automatically closed in LineDecoder.
	decoder := bulk.LineDecoder{Body: file}
	progress := make(chan bulk.Status)

	// Feed the URLs with a decoder.
	bulker.Feed(decoder, progress)
	for status := range progress {
		fmt.Println(status, status.Redirs)
	}
}
