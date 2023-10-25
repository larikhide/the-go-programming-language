package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Fetch prints the content found at each specified URL.
func Fetch1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %v]n", err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}

// 1.7
func Fetch2() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %v]n", err)
			os.Exit(1)
		}
	}
}

// 1.8
func Fetch3() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %v]n", err)
			os.Exit(1)
		}
	}
}

// 1.9
func Fetch4() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, strings.NewReader(resp.Status+"\n"))
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %v]n", err)
			os.Exit(1)
		}
	}
}
