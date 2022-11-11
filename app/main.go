package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", simpleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func simpleHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello world!")
}
