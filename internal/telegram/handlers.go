package telegram

import (
	"gopkg.in/telebot.v3"
)

func (t *Telegram) setupHandlers() {
	// middelwares
	t.bot.Use(t.registerMiddelware)

	// handlers
	t.bot.Handle("/start", t.start)
	t.bot.Handle(telebot.OnText, t.textHandler)
	t.bot.Handle(&brnEditDisplayName, t.editDisplayName)
}

func (t *Telegram) textHandler(c telebot.Context) error {
	if t.TelePrompt.Dispatch(c.Sender().ID, c) {
		return nil
	}
	return c.Reply("I can not handle this message")
}
