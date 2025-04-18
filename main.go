package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, req *http.Request) {
	log.Print("hello-world: received a request")
	fmt.Fprintf(writer, "Hello World!\n")
}

func main() {
	log.Print("hello-world: starting server...")
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("hello-world: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
