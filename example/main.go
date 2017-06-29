package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"../../bulk"
)

/*
	Author: 	Adrian Plavka
	Contact: 	adrian.plavka@gmail.com
	Reference: 	https://www.github.com/adrianplavka

	README!
	This is a console application named "bulk".
	It will run a batch of URLs and provide information, if they:
	are Valid and where they Redirect, or Invalid.

	Every URL checking is done by a goroutine (concurrently).
*/

func main() {
	// Bulker is a type, that is defined by http.Client.
	bulker := bulk.DefaultBulker

	// Open a CSV file for read-only.
	path, _ := filepath.Abs("bulk/example/urls.csv")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("failed while opening a file: ", err)
	}

	decoder := bulk.LineDecoder{Body: file}
	progress := make(chan bulk.Status)

	bulker.Feed(decoder, progress)
	for status := range progress {
		fmt.Println(status)
	}
}
