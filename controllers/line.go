package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/yusuke9929/yotawa-with-go/helpers"
	"github.com/yusuke9929/yotawa-with-go/models"
)

// added session to our lineController
type LineController struct {
	session            string
	channelSecret      string
	channelAccessToken string
}

// added session to our lineController
func NewLineController(s string) *LineController {
	return &LineController{
		"session",
		os.Getenv("channelSecret"),
		os.Getenv("channelAccessToken"),
	}
}

func (lc LineController) Callback(w http.ResponseWriter, req *http.Request) {

	bot, err := linebot.New(lc.channelSecret, lc.channelAccessToken)
	helpers.CheckError(err)

	events, err := bot.ParseRequest(req)
	helpers.CheckError(err)

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
		var MessagesToReply []linebot.Message
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				ContentsToReply := models.GetAutoReplyContents(message.Text)
				MessagesToReply = helpers.ConvertContentsToMessages(ContentsToReply)
			case *linebot.ImageMessage:
				// Not Yet Implementes
			}
		}
		// execute message-reply
		if _, err = bot.ReplyMessage(event.ReplyToken, MessagesToReply...).Do(); err != nil {
			log.Print(err)
		}

	}
}
