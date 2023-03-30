package main

import (
	"fmt"
	"net/http"
)

const (
	listenAddr = "0.0.0.0:8080"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	fmt.Printf("Starting server on %s\n", listenAddr)
	http.ListenAndServe(listenAddr, nil)
}

func helloWorldHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello World"))
}
