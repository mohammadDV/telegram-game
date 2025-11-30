package telegram

import (
	"errors"

	"github.com/samber/lo"
	"gopkg.in/telebot.v3"
)

var (
	ErrorInputTimeout  = errors.New("input timeout")
	ErrorInputCanceled = errors.New("input canceled")
)

type Confirm struct {
	ConfirmText func(msg *telebot.Message) string
}

type Validator struct {
	validator func(msg *telebot.Message) bool
	OnInvalid func(msg *telebot.Message) string
}

type InputContext struct {
	Prompt         any
	PromptKeyboard [][]string

	OnTimeout any
	Validator Validator
	Confirm   Confirm
}

func (t *Telegram) Input(c telebot.Context, config InputContext) (*telebot.Message, error) {
getInput:
	if config.Prompt != nil {
		if config.PromptKeyboard != nil {
			c.Send(config.Prompt, generateKeyboard(config.PromptKeyboard))
		} else {
			c.Send(config.Prompt)
		}
	}

	response, isTimeout := t.TelePrompt.AsMessage(c.Sender().ID, DefaultInputTimeout)
	if isTimeout {
		if config.OnTimeout != nil {
			c.Send(config.OnTimeout)
		} else {
			c.Send(DefaultTimeoutText)
		}
		return nil, ErrorInputTimeout
	}

	if config.Validator.validator != nil && !config.Validator.validator(response) {
		if !config.Validator.validator(response) {
			c.Send(config.Validator.OnInvalid(response))
			goto getInput
		}
	}

	// client has to confirm
	if config.Confirm.ConfirmText != nil {
		confirmText := config.Confirm.ConfirmText(response)
		confirmResponse, err := t.Input(c, InputContext{
			Prompt:         confirmText,
			PromptKeyboard: [][]string{{ConfirmText, DeclineText}},
			Validator:      chooseValidator(ConfirmText, DeclineText),
		})
		if err != nil {
			return nil, err
		}
		if confirmResponse.Text == DeclineText {
			goto getInput
		}
		// if confirmResponse.Text == ConfirmText {
		// 	return response, nil
		// }
		// return nil, ErrorInputCanceled
	}

	return response, nil
}

func generateKeyboard(rows [][]string) *telebot.ReplyMarkup {
	mu := &telebot.ReplyMarkup{ResizeKeyboard: true, OneTimeKeyboard: true}
	replyRows := lo.Map(rows, func(row []string, _ int) telebot.Row {
		btns := lo.Map(row, func(item string, _ int) telebot.Btn {
			return mu.Text(item)
		})
		return mu.Row(btns...)
	})
	mu.Reply(replyRows...)
	return mu
}
