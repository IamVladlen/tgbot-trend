package tgbot

import (
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type bot struct {
	bot     *telego.Bot
	Handler *th.BotHandler
}

func New(token string) *bot {
	b, err := telego.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	updates, err := b.UpdatesViaLongPolling(nil)
	if err != nil {
		log.Fatal(err)
	}

	h, err := th.NewBotHandler(b, updates)
	if err != nil {
		log.Fatal(err)
	}

	return &bot{
		bot:     b,
		Handler: h,
	}
}

func (b *bot) Start() {
	go func() {
		fmt.Println("Bot has successfully launched")
		b.Handler.Start()
	}()
}

func (b *bot) Stop() {
	b.bot.StopLongPolling()
	b.Handler.Stop()
}
