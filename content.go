package main

import "strings"

type content struct {
	Ctype  string
	Ctext  string
	Cimage string
	Clink  string
}

func newContentText(text string) content {
	return content{
		Ctype: "text",
		Ctext: text,
	}
}

func newContentImage(text string, image string) content {
	return content{
		Ctype:  "image",
		Ctext:  text,
		Cimage: image,
	}
}

func newContentLink(text string, image string, link string) content {
	return content{
		Ctype:  "link",
		Ctext:  text,
		Cimage: image,
		Clink:  link,
	}
}

func getAutoReplyContents(s string) []content {
	var replyMessages []content
	for key, value := range dictionary {
		if strings.Contains(s, key) {
			replyMessages = append(replyMessages, newContentText(value))
		}
	}
	return replyMessages
}
