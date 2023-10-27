package ch1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func Server1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request URL.Path=%q\n", r.URL.Path)
}

var mu sync.Mutex
var count int

func Server2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count %d\n", count)
	mu.Unlock()
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}

	fmt.Fprintf(w, "Host=%q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
	}
}

func handlerLissajous(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	cycleStr := r.FormValue("cycles")
	if cycleStr != "" {
		cycles, _ = strconv.Atoi(cycleStr) //ignoring err
	}
	LissajousForServer(w, cycles)

}

func Server3() {
	http.HandleFunc("/", handler3)
	http.HandleFunc("/count", counter2)
	http.HandleFunc("/lissajous", handlerLissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
