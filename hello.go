package main

import (
	"fmt"
	"net/http"
)

func httpHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, world!")
}

func main() {
	http.HandleFunc("/", httpHandler)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
