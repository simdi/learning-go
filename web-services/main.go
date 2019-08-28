package main

import (
	"fmt"
	"net/http"
)

func handleRequest(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	fmt.Fprintf(response, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8100", nil)
}
