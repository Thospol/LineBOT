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
	client             = &http.Client{}
	channelSecret      = "7e602176049bb7bad6058a89cf0a4df1"
	channelAccsssToken = "gg35kI4YEqjkxhxyiMTlU6Mx7sFJHwi5iR5WPKV2mshDnUAJzQjWZwzknFnNRVOUDIl91oJGxV7SLivTSmRNW9jqz1fVtAEwC4Y7sFfhj2rB9qqYWMdvXQOq6L03ceChaAcgI+o+zUNT8ae8nxEvPwdB04t89/1O/w1cDnyilFU="
)

func main() {

	var err error
	bot, err = linebot.New(channelSecret, channelAccsssToken)
	log.Println("Bot:", bot, " err:", err)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/callback", handler)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
