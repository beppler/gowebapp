package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", hello)

	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		panic(err)
	}
}
