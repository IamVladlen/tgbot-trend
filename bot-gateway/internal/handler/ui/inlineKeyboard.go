package ui

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	_maxRows = 2
)

func InlineButton(cmd string) *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(cmd).WithCallbackData(cmd),
		),
	)
}

func InlineButtons(cmd1, cmd2, cmd3 string) *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(cmd1).WithCallbackData(cmd1),
			tu.InlineKeyboardButton(cmd2).WithCallbackData(cmd2),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(cmd3).WithCallbackData(cmd3),
		),
	)
}

func InlineButtonsSchedule(cmd1, cmd2, cmd3 string) *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(cmd1).WithCallbackData(cmd1),
			tu.InlineKeyboardButton(cmd2).WithCallbackData(cmd2),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(cmd3).WithCallbackData(cmd3),
		),
	)
}

func InlineCountries(countries []string) *telego.InlineKeyboardMarkup {
	rows := make([][]telego.InlineKeyboardButton, _maxRows)
	for i := range rows {
		rows[i] = make([]telego.InlineKeyboardButton, len(countries)/len(rows))
	}
	var lastIdx int

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			rows[i][j] = tu.InlineKeyboardButton(countries[lastIdx]).WithCallbackData(countries[lastIdx])
			lastIdx++
		}
	}

	return tu.InlineKeyboard(
		rows...,
	)
}
