package bulk

import (
	"net/http"
	"strconv"
	"strings"
)

// URL valid status messages.
const (
	StatusValid   = "URL <{}> is Valid. "
	StatusInvalid = "URL <{}> is Invalid. "
)

// redirection is a struct that consists of a redirected URL
// and a StatusCode.
type redirection struct {
	URL        string
	StatusCode int
}

func (r redirection) String() string {
	return strconv.Itoa(r.StatusCode) + " " + http.StatusText(r.StatusCode) + " - " + r.URL
}

// Status is a struct that indicates if an URL was valid,
// and how many redirects it had.
type Status struct {
	URL    string
	Valid  bool
	Redirs []redirection
}

// Function handleRedirection represents a way to count the amount of redirections
// made by the HTTP request.
// This overrides a CheckRedirect function in http.Client.
func (s *Status) handleRedirection(req *http.Request, via []*http.Request) error {
	redir := redirection{req.URL.String(), req.Response.StatusCode}
	s.Redirs = append(s.Redirs, redir)
	return nil
}

func (s Status) String() (msg string) {
	if s.Valid {
		msg = strings.Replace(StatusValid, "{}", s.URL, -1)
	} else {
		msg = strings.Replace(StatusInvalid, "{}", s.URL, -1)
	}
	return
}
