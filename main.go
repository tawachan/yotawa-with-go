package main

import (
	"io"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func main() {

	// lc := controllers.NewLineController(getSession())

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", root)
	// http.HandleFunc("/callback", lc.Callback)

	http.ListenAndServe(":"+port, nil)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

//Route for Checking whether application is running
func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Your app is now runninng !!!")
}
