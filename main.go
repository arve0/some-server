package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8090")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from some-server\n")
}
