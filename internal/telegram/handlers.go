package telegram

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func (t *Telegram) setupHandlers() {
	// middelwares
	t.bot.Use(t.registerMiddelware)

	// handlers
	t.bot.Handle("/start", t.start)
	t.bot.Handle(telebot.OnText, func(c telebot.Context) error {
		if t.TelePrompt.Dispatch(c.Sender().ID, c) {
			return nil
		}
		return c.Reply("I can not handle this message")
	})
}

func (t *Telegram) start(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)
	_ = isJustCreated

	msg, err := t.Input(c, InputContext{
		Prompt:    "Welcome to the game! please enter your name",
		OnTimeout: "Timeout too",
	})
	if err != nil {
		return err
	}
	c.Reply(fmt.Sprintf("Hello, %s!", msg.Text))

	return nil
}
