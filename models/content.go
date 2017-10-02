package models

import "strings"

type Content struct {
	Ctype  string
	Ctext  string
	Cimage string
	Clink  string
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
		Ctype: "text",
		Ctext: text,
	}
}

func NewContentImage(text string, image string) Content {
	return Content{
		Ctype:  "image",
		Ctext:  text,
		Cimage: image,
	}
}

func NewContentLink(text string, image string, link string) Content {
	return Content{
		Ctype:  "link",
		Ctext:  text,
		Cimage: image,
		Clink:  link,
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
