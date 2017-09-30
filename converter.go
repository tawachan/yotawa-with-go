package main

import "github.com/line/line-bot-sdk-go/linebot"

func convertContentsToMessages(contents []content) (messagesToReply []linebot.Message) {
	// Set up contents for line format
	for _, c := range contents {
		var lm linebot.Message
		if c.Ctype == "text" {
			lm = linebot.NewTextMessage(c.Ctext)
		} else if c.Ctype == "image" {
			lm = linebot.NewImageMessage(c.Cimage, c.Cimage)
		} else if c.Ctype == "link" {
			lm = makeMessageWithCarousel(c)
		} else {
			continue
		}
		messagesToReply = append(messagesToReply, lm)
	}
	return messagesToReply
}

func makeMessageWithCarousel(c content) linebot.Message {
	title := c.Ctext
	link := c.Clink
	desc := "Check this out!"
	image := "https://68.media.tumblr.com/7433692cabbfa132f34adb034e7909fa/tumblr_inline_owu4b4v7ow1v9napg_500.png"

	action := linebot.NewURITemplateAction("View", link)
	carousel := linebot.NewCarouselColumn(image, title, desc, action)
	template := linebot.NewCarouselTemplate(carousel)

	return linebot.NewTemplateMessage("Check this out!", template)
}
