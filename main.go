package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	bot                *linebot.Client
	channelSecret      = "198e00458c411f2c50aa19beaa94851d"
	channelAccsssToken = "2A3RczSVO44rsUJvAdd8utWfhOyELWeMVnIQ5IVDC6A9sCjSxFhcf5K1u/KuanMNDIl91oJGxV7SLivTSmRNW9jqz1fVtAEwC4Y7sFfhj2rHioE7uHVJaXDny55T+WXYOV4qQ2gubjAG7rsIVFUUTAdB04t89/1O/w1cDnyilFU="
)
var recentId string

func main() {

	var err error
	bot, err = linebot.New(channelSecret, channelAccsssToken)
	log.Println("Bot:", bot, " err:", err)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/push", pushMessage)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		fmt.Println(event.Source.UserID)
		recentId = event.Source.UserID
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

func pushMessage(w http.ResponseWriter, r *http.Request) {
	msg := linebot.NewTextMessage("Test")
	bot.PushMessage(recentId, msg)
}
