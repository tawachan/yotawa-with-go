package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/callback", callback)

	http.ListenAndServe(":"+port, nil)

}

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "test")
}

func callback(w http.ResponseWriter, req *http.Request) {
	channelSecret := os.Getenv("channelSecret")
	channelAccessToken := os.Getenv("channelAccessToken")
	io.WriteString(w, "callback")
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
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
