package bulk

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// Bulker type, that implements an HTTP client.
type Bulker struct {
	client http.Client
}

// DefaultBulker is a default implementation of Bulker.
// The timeout is a maximum of 5 seconds
// and the amount of redirection it can have is 10 (http.Client default).
var DefaultBulker = Bulker{
	http.Client{
		Timeout: time.Duration(time.Second * 5),
	},
}

// Check checks an URL from a raw url string.
func (b Bulker) Check(url string) Status {
	s := Status{URL: url} // Create a URL status.
	// If any redirection happens, the redirected URLs will be added to the status.
	b.client.CheckRedirect = s.handleRedirection

	// Actual HEAD request to the URL, ignoring entire body.
	// If everything went OK, we treat an URL as Valid.
	// If there was some kind of error, we treat an URL as Invalid.
	req, _ := http.NewRequest("HEAD", url, nil)
	if _, err := b.client.Do(req); err == nil {
		s.Valid = true
	} else {
		s.Valid = false
	}
	return s
}

// CheckMultiple is a concurrent function, that checks multiple URLs.
func (b Bulker) CheckMultiple(urls []string, status chan<- Status) {
	// Main goroutine, which waits for all URLs to be checked
	// and sends status to the channel.
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(urls))
		for _, url := range urls {
			url := url
			go func() {
				defer wg.Done()
				status <- b.Check(url)
			}()
		}
		wg.Wait()
		close(status)
	}()
}

// Feed requires a Decoder, from which body it will decode URLs
// and then reports the information to the Status channel.
func (b Bulker) Feed(d Decoder, s chan<- Status) {
	urls, err := d.Decode()
	if err != nil {
		log.Fatalln("error during decoding: ", err)
	}
	b.CheckMultiple(urls, s)
}
