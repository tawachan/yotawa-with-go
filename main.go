package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/callback", callback)

	http.ListenAndServe(":"+port, nil)
}
