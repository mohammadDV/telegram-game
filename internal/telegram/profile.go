package telegram

import (
	"context"
	"fmt"

	"gopkg.in/telebot.v3"
)

func (t *Telegram) editDisplayName(c telebot.Context) error {
	c.Delete()
	return t.editDisplayNameProfile(c, "What name do you want to use?")
}

func (t *Telegram) editDisplayNameProfile(c telebot.Context, promptText string) error {

	account := getAccount(c)
	msg, err := t.Input(c, InputContext{
		Prompt:    promptText,
		OnTimeout: "Timeout too",
		Validator: Validator{
			validator: func(msg *telebot.Message) bool {
				len := len([]rune(msg.Text))
				return len >= 3 && len <= 20
			},
			OnInvalid: func(msg *telebot.Message) string {
				return "Your name should be between 3 and 20 characters"
			},
		},
		Confirm: Confirm{
			ConfirmText: func(msg *telebot.Message) string {
				return fmt.Sprintf("Hello, %s! Are you sure you want to use this name?", msg.Text)
			},
			// DeclineText: func(msg *telebot.Message) string {
			// 	return fmt.Sprintf("Hello, %s! Are you sure you want to use this name?", msg.Text)
			// },
		},
	})
	if err != nil {
		return err
	}

	displayname := msg.Text
	account.DisplayName = displayname

	if err := t.App.Account().Update(context.Background(), account); err != nil {
		return err
	}
	c.Set("account", account)

	// todo: validation
	c.Reply(fmt.Sprintf("Hello, %s!", displayname))

	return nil
}
