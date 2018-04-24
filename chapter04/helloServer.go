package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldhandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Received request [%s] for path: [%s]", r.Method, r.URL.Path)
	log.Println(msg)

	response := fmt.Sprintf("Hello, World! at Path: %s", r.URL.Path)
	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/", helloWorldhandler) // Catch all path
	log.Println("Starting server at port: 8080...")
	http.ListenAndServe(":8080", nil)
}
