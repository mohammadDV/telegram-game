package telegram

import (
	"context"
	"fmt"
	"gopkg.in/telebot.v3"
)

func (t *Telegram) start(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)
	account := getAccount(c)

	if isJustCreated {
		//todo: ...
	}

	msg, err := t.Input(c, InputContext{
		Prompt:    "Welcome to the game! please enter your name",
		OnTimeout: "Timeout too",
	})
	if err != nil {
		return err
	}

	displayname := msg.Text
	account.DisplayName = displayname 

	if err :=  t.App.Account().Update(context.Background(), account); err != nil {
		return err
	}

	// todo: validation
	c.Reply(fmt.Sprintf("Hello, %s!", displayname))

	return nil
}