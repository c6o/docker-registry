package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from K8S")
	})

	log.Println("Starting Hello on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error starting up server", err)
	}
}
