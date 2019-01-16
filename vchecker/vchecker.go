package vchecker

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://go.googlesource.com/go/+/go"
)

// Version checks if the given Go version exists
func Version(v string, client *http.Client) (bool, error) {
	if client == nil {
		client = &http.Client{
			Timeout: 5 * time.Second,
		}
	}

	r, err := client.Head(fmt.Sprintf("%s%s", baseURL, v))

	if err != nil {
		return false, err
	}

	if r.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}

// Output outputs version result
func Output(w io.Writer, exists bool, err error) {
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Some error occurred: %s\n", err.Error())))
		return
	}

	if !exists {
		w.Write([]byte("It does not exist\n"))
		return
	}

	w.Write([]byte("It exists!\n"))
}
