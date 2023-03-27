package tgbot

import (
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type bot struct {
	Bot     *telego.Bot
	Handler *th.BotHandler
}

// New creates bot instance with handler.
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
		Bot:     b,
		Handler: h,
	}
}

// Start starts updates handling.
func (b *bot) Start() {
	go func() {
		fmt.Println("Bot has successfully launched")
		b.Handler.Start()
	}()
}

// Stop gracefully stops updates receiving.
func (b *bot) Stop() {
	b.Bot.StopLongPolling()
	b.Handler.Stop()
}
