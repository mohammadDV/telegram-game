package telegram

import (
	"errors"
	"gopkg.in/telebot.v3"
)

var (
	ErrorInputTimeout  = errors.New("input timeout")
	ErrorInputCanceled = errors.New("input canceled")
)

type InputContext struct {
	Prompt    any
	OnTimeout any
}

func (t *Telegram) Input(c telebot.Context, config InputContext) (*telebot.Message, error) {

	if config.Prompt != nil {
		c.Reply(config.Prompt)
	}

	response, isTimeout := t.TelePrompt.AsMessage(c.Sender().ID, DefaultInputTimeout)
	if isTimeout {
		if config.OnTimeout != nil {
			c.Reply(config.OnTimeout)
		} else {
			c.Reply(DefaultTimeoutText)
		}
		return nil, ErrorInputTimeout
	}
	return response, nil
}