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
	io.WriteString(w, "Your App is now runninng !!!")
}

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
		var replyTexts []string
		var replyMessages []linebot.Message
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyTexts = getAutoResponses(message.Text)

			case *linebot.ImageMessage:
				replyTexts = append(replyTexts, "image included")
			}
		}

		// Set line structs, using slices set up above
		for _, m := range replyTexts {
			replyMessages = append(replyMessages, linebot.NewTextMessage(m))
		}

		// execute message-reply
		if _, err = bot.ReplyMessage(event.ReplyToken, replyMessages...).Do(); err != nil {
			log.Print(err)
		}

	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
