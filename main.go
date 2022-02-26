package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	fmt.Printf("Starting server on port %s\n", port)
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("request from %s for %s\n", req.RemoteAddr, req.Header.Get("x-forwarded-for"))
	fmt.Fprintf(w, "some-server\n")
}
