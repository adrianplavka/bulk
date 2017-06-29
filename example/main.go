package main

import (
	"fmt"

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

	urls := []string{
		"http://google.com",
		"http://github.com",
		"http://stackoverflow.com"}
	progress := make(chan bulk.Status)

	bulker.CheckMultiple(progress, urls[:]...)
	for status := range progress {
		fmt.Println(status)
	}
}
