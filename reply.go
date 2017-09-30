package main

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

type dict map[string]string
type content struct {
	Ctype   string
	Content string
}

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

func getReplyContents(t string) (replyContents []linebot.Message) {
	var contents []content
	contents = getAutoResponses(t)
	if len(contents) == 0 {
		contents = append(contents, content{"link", t})
	}
	replyContents = convertToLineFormat(contents)
	return
}

func convertToLineFormat(contents []content) (replyContents []linebot.Message) {
	// Set up contents for line format
	for _, c := range contents {
		var rc linebot.Message
		if c.Ctype == "text" {
			rc = linebot.NewTextMessage(c.Content)
		} else if c.Ctype == "image" {
			rc = linebot.NewImageMessage(c.Content, c.Content)
		} else if c.Ctype == "link" {
			title := "Tumblr"
			link := c.Content
			desc := "Check this out!"
			image := "https://68.media.tumblr.com/7433692cabbfa132f34adb034e7909fa/tumblr_inline_owu4b4v7ow1v9napg_500.png"

			action := linebot.NewURITemplateAction("View", link)
			carousel := linebot.NewCarouselColumn(image, title, desc, action)
			template := linebot.NewCarouselTemplate(carousel)

			rc = linebot.NewTemplateMessage("Check this out!", template)
		} else {
			continue
		}
		replyContents = append(replyContents, rc)
	}
	return replyContents
}

func getAutoResponses(s string) []content {
	var replyMessages []content
	for key, value := range dictionary {
		if strings.Contains(s, key) {
			replyMessages = append(replyMessages, content{"text", value})
		}
	}
	return replyMessages
}

func getResponceDict() dict {
	return dictionary
}
