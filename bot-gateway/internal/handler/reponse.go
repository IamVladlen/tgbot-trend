package handler

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func response(bot *telego.Bot, id int64, markup telego.ReplyMarkup, msg string) error {
	m := tu.Message(
		tu.ID(id),
		msg,
	).WithReplyMarkup(markup)

	_, err := bot.SendMessage(m)
	if err != nil {
		return err
	}

	return nil
}

func responseEntities(bot *telego.Bot, id int64, markup telego.ReplyMarkup, msg ...tu.MessageEntityCollection) error {
	m := tu.MessageWithEntities(
		tu.ID(id),
		msg...,
	).WithReplyMarkup(markup)

	_, err := bot.SendMessage(m)
	if err != nil {
		return err
	}

	return nil
}
