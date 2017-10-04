package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/yusuke9929/yotawa-with-go/controllers"
)

func main() {
	lc := controllers.NewLineController()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/callback", lc.Callback)

	http.ListenAndServe(":"+port, nil)
}

//Route for Checking whether application is running
func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Your app is now runninng !!!")
}
