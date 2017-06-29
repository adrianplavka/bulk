package bulk

import (
	"encoding/csv"
	"io"
)

// Decoder is an interface, that should decode a file and return the URLs.
// If needed, implement your own type, that returns a string slice of URLs,
// and an error.
type Decoder interface {
	Decode() ([]string, error)
}

// LineDecoder represents a line-by-line body, from which the URLs are being decoded from.
type LineDecoder struct {
	Body io.ReadCloser
}

// Decode a body, that separates each URL with a newline.
func (d LineDecoder) Decode() ([]string, error) {
	defer d.Body.Close()
	csvr := csv.NewReader(d.Body) // Create a new CSV reader.

	var urls []string
	for { // Read a record line by line.
		record, err := csvr.Read()
		if err == io.EOF {
			return urls, nil
		} else if err != nil {
			return nil, err
		}
		urls = append(urls, record[0])
	}
}
