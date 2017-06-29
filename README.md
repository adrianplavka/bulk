# Bulk

URL checking library for Go.

This library helps to provide a simple API for testing *multiple URLs* for their validness -
if they are **Valid** and where they **__Redirect__**, or if they are **Invalid**.
Requests are done by a **"HEAD" HTTP method**, and every URL checking is done **concurrently**.

Maintaned by @adrianplavka.

## Installation

```
go get -u github.com/adrianplavka/bulk
```

## Usage

For simple usage, simply declare a DefaultBulker:

```go
import "github.com/adrianplavka/bulk"

func main() {
    bulker := bulk.DefaultBulker
```

You can now use your Bulker, which has methods pre-defined.

To check a single URL for it's validness, simply:

```go
    status := bulker.Check("http://www.google.com")
    fmt.Println(status)
```

Check method returns a single Status, that simply tells if the URL was Valid with Redirections or Not.

```go
type Status struct {
	URL    string
	Valid  bool
	Redirs []redirection
}
```

To check how many redirections it had, you can iterate over status' Redirs:

```go
    fmt.Println(status.Redirs)
```

To check multiple URLs, you pass a string slice and a status channel to CheckMultiple method:

```go
    progress := make(chan bulk.Status)

    bulker.CheckMultiple(urls, progress)
    for status := range progress {
        fmt.Println(status)
    }
```

This loop blocks until every URL has been checked.

To check URLs from a file or a body, Feed method expects a Decoder interface wih a Status channel.
Decoder interface contains only one method Decode, where you specify how you obtain URLs from a source.
Decode should return a string slice of URLs and an error, if occured.

```go
type Decoder interface {
	Decode() ([]string, error)
}
```

Bulk comes with a LineDecoder, that reads a source line-by-line, URLs are separated by a newline.

```go
type LineDecoder struct {
	Body io.ReadCloser
}
```

To use this, simply declare a LineDecoder with a Body (that is automatically closed after decoding).

```go
    // Open a CSV file for read-only.
	path := "urls.csv"
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
```

## License

**MIT**