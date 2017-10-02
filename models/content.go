package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Content struct {
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

func GetAutoReplyContents(s string) []Content {
	var contents []Content
	for key, value := range dictionary {
		if strings.Contains(s, key) {
			contents = append(contents, NewContentText(value))
		}
	}
	if len(contents) == 0 {
		contents = append(contents, NewContentText(s))
	}
	return contents
}

func (c Content) getContents(s string) []Content {
	db, err := sql.Open("postgres", "user=postgres dbname=yotawa-with-go")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM contents where key like $1", s)
	fmt.Println(rows.Scan())

	var userid int
	err = db.QueryRow(`INSERT INTO users(name, age)
		VALUES('beatrice', 93) RETURNING id`).Scan(&userid)

	fmt.Println(err)
	fmt.Println(userid)
	return []Content{}
}
