package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server listening at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
