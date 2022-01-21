package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8090")
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8090", nil)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "some-server\n")
}
