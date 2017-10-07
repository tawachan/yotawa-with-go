package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/yotawa9929/yotawa-with-go/logs"
)

type Content struct {
	Id       string
	Category string
	Key      string
	Text     string
	Image    string
	Link     string
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
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	logs.CheckError(err)

	rows, err := db.Query("SELECT * FROM contents")
	defer rows.Close()

	for rows.Next() {
		var c Content
		if err := rows.Scan(&c.Id, &c.Category, &c.Key, &c.Text, &c.Image, &c.Link); err != nil {
			log.Fatal(err)
		}
		fmt.Println(c)
		if strings.Contains(s, c.Key) {
			contents = append(contents, c)
		}
	}

	if len(contents) == 0 {
		contents = append(contents, NewContentText(s))
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}
