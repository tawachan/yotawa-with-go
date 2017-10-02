package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Content struct {
	Id       string
	Category string
	Key      string
	Text     string
	Image    string
	Link     string
}

// Will be deleted when mongo db connected
type dict map[string]string

var dictionary dict

func init() {
	dictionary = dict{
		"おおたわ":   "おおたわです",
		"OTYM":   "おたやま",
		"test":   "test desu",
		"コンパイル":  "難しい",
		"tumblr": "こちらを見てね→ https://yotawa9929.tumblr.com/",
	}
}

func NewContentText(text string) Content {
	return Content{
		Category: "text",
		Text:     text,
	}
}

func NewContentImage(text string, image string) Content {
	return Content{
		Category: "image",
		Text:     text,
		Image:    image,
	}
}

func NewContentLink(text string, image string, link string) Content {
	return Content{
		Category: "link",
		Text:     text,
		Image:    image,
		Link:     link,
	}
}

// func GetAutoReplyContents(s string) []Content {
// 	var contents []Content
// 	for key, value := range dictionary {
// 		if strings.Contains(s, key) {
// 			contents = append(contents, NewContentText(value))
// 		}
// 	}
// 	if len(contents) == 0 {
// 		contents = append(contents, NewContentText(s))
// 	}
// 	return contents
// }

func GetAutoReplyContents(s string) []Content {
	var contents []Content
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT * FROM contents where key like $1", s)
	defer rows.Close()

	for rows.Next() {
		var c Content
		if err := rows.Scan(&c.Id, &c.Category, &c.Key, &c.Text, &c.Image, &c.Link); err != nil {
			log.Fatal(err)
		}
		fmt.Println(c)
		contents = append(contents, c)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}
