package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello world from %s/%s!\n", runtime.GOOS, runtime.GOARCH))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	fmt.Println("Listening on port 8000")
	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		panic(err)
	}
}
