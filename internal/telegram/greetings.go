package telegram

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func (t *Telegram) start(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)

	if !isJustCreated {
		return t.myInfo(c)
	}

	if err := t.editDisplayNameProfile(c, "Welcome to the game! please enter your name"); err != nil {
		return err
	}

	return t.myInfo(c)
}


func (t *Telegram) myInfo(c telebot.Context) error {
	account := getAccount(c)

	var selector = &telebot.ReplyMarkup{}
	selector.Inline(selector.Row(brnEditDisplayName))
	return c.Send(fmt.Sprintf("Welcome, how can I help you? %s!", account.DisplayName),
		selector,
)

}