package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/request/", requestHandler)
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
