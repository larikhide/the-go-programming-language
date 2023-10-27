package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Fetchall fetches URLs in parallel and reports their times and sizes.
func Fetchall() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err, "\n")
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

// FetchallToFile fetches URLs in parallel and reports their times and sizes.
func FetchallToFile() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	file, err := os.OpenFile("fetchall.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		ch <- fmt.Sprintf("while creating a file: %v\n", err)
		return
	}
	defer file.Close()

	for range os.Args[1:] {
		data := <-ch
		_, err := file.WriteString(data)
		if err != nil {
			ch <- fmt.Sprintf("while wreating in file %s: %v\n", data, err)
			continue
		}
		file.Sync()
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
