package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

//Route for Checking whether application is running
func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Your app is now runninng !!!")
}

//Route for Linebot
func callback(w http.ResponseWriter, req *http.Request) {
	channelSecret := os.Getenv("channelSecret")
	channelAccessToken := os.Getenv("channelAccessToken")

	bot, err := linebot.New(channelSecret, channelAccessToken)
	checkError(err)

	events, err := bot.ParseRequest(req)
	checkError(err)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	// Set up slices for reply
	for _, event := range events {
		var replyContents []linebot.Message
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyContents = getReplyContents(message.Text)
			case *linebot.ImageMessage:
				replyContents = convertToLineFormat([]content{content{"text", "image detected"}})
			}
		}

		// execute message-reply
		if _, err = bot.ReplyMessage(event.ReplyToken, replyContents...).Do(); err != nil {
			log.Print(err)
		}

	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
