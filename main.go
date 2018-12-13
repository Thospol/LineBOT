package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
)

func main() {
	handler, err := httphandler.New("7e602176049bb7bad6058a89cf0a4df1", "VJyEo8k4B1hD5Ql/b3/lVloeL0yAiOW28Hpjw/BSFjwXAMryXYqs/NLLB+qjDVT+DIl91oJGxV7SLivTSmRNW9jqz1fVtAEwC4Y7sFfhj2q5h2vOBteXaRr4nDVk7kEy+seAaM4iFnU4sBaSmG2lEAdB04t89/1O/w1cDnyilFU=")
	log.Println("Bot:", handler, " err:", err)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
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
	})
	http.Handle("/callback", handler)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
