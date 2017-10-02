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
	desc := "Check this out!"
	image := "https://68.media.tumblr.com/7433692cabbfa132f34adb034e7909fa/tumblr_inline_owu4b4v7ow1v9napg_500.png"

	action := linebot.NewURITemplateAction("View", link)
	carousel := linebot.NewCarouselColumn(image, title, desc, action)
	template := linebot.NewCarouselTemplate(carousel)

	return linebot.NewTemplateMessage("Check this out!", template)
}
