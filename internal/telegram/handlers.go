package telegram

import (
	"gopkg.in/telebot.v3"
)

func (t *Telegram) setupHandlers() {
	t.bot.Use(t.registerMiddelware)
	t.bot.Handle("/start", t.start)
}

func (t *Telegram) start(c telebot.Context) error {
	return c.Reply("Hello, world!")
}