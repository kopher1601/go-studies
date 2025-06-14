package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchRemoteResource(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	return io.ReadAll(r.Body)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specifyt a HTTp URL to get data from")
		os.Exit(1)
	}

	body, err := fetchRemoteResource(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error fetching remote resource: %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s", body)
}
