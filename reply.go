package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

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

func getMessagesToText(t string) (messagesToReply []linebot.Message) {
	var contents []content
	contents = getAutoReplyContents(t)
	if len(contents) == 0 {
		contents = append(contents, newContentText(t))
	}
	messagesToReply = convertContentsToMessages(contents)
	return
}

func getMessagesToImage() (messagesToReply []linebot.Message) {
	messagesToReply = convertContentsToMessages([]content{newContentText("image detected")})
	return
}
