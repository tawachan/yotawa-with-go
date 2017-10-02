package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/yusuke9929/yotawa-with-go/controllers"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=yotawa-with-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM users where id = 1")
	fmt.Println(rows.Scan())

	var userid int
	err = db.QueryRow(`INSERT INTO users(name, age)
		VALUES('beatrice', 93) RETURNING id`).Scan(&userid)

	fmt.Println(err)
	fmt.Println(userid)
	lc := controllers.NewLineController(getSession())

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/callback", lc.Callback)

	http.ListenAndServe(":"+port, nil)
}

func getSession() string {
	return "session"
}

//Route for Checking whether application is running
func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Your app is now runninng !!!")
}
