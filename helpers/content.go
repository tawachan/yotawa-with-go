package helpers

import "github.com/line/line-bot-sdk-go/linebot"
import "github.com/yusuke9929/yotawa-with-go/models"

func ConvertContentsToMessages(contents []models.Content) (messagesToReply []linebot.Message) {
	// Set up contents for line format
	for _, c := range contents {
		var lm linebot.Message
		if c.Category == "text" {
			lm = MakeMessageWithText(c)
		} else if c.Category == "image" {
			lm = MakeMessageWithImage(c)
		} else if c.Category == "link" {
			lm = MakeMessageWithCarousel(c)
		} else {
			continue
		}
		messagesToReply = append(messagesToReply, lm)
	}
	return messagesToReply
}

func MakeMessageWithText(c models.Content) linebot.Message {
	return linebot.NewTextMessage(c.Text)
}

func MakeMessageWithImage(c models.Content) linebot.Message {
	return linebot.NewImageMessage(c.Image, c.Image)
}

func MakeMessageWithCarousel(c models.Content) linebot.Message {
	title := c.Text
	link := c.Link
	desc := c.Link
	image := c.Image

	action := linebot.NewURITemplateAction("View", link)
	carousel := linebot.NewCarouselColumn(image, title, desc, action)
	template := linebot.NewCarouselTemplate(carousel)

	return linebot.NewTemplateMessage(desc, template)
}
