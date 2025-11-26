package TelePrompt

import (
	"sync"

	"gopkg.in/telebot.v3"
	"time"
)

type Prompt struct {
	TeleCtx telebot.Context
}

type TelePrompt struct {
	accountPrompts sync.Map
}

func NewTelePrompt() *TelePrompt {
	return &TelePrompt{}
}

func (t *TelePrompt) Register(userId int64) <-chan Prompt {

	ch := make(chan Prompt, 1)

	if preChannel, loaded := t.accountPrompts.LoadAndDelete(userId); loaded {
		close(preChannel.(chan Prompt))
	}

	t.accountPrompts.Store(userId, ch)

	return ch
}

func (t *TelePrompt) AsMessage(userId int64, timeout time.Duration) (*telebot.Message, bool) {
	ch := t.Register(userId)
	select {
	case val := <-ch:
		return val.TeleCtx.Message(), false
	case <-time.After(timeout):
		return nil, true
	}
}

func (t *TelePrompt) Dispatch(userId int64, ctx telebot.Context) bool {
	ch, loaded := t.accountPrompts.LoadAndDelete(userId)
	if !loaded {
		return false
	}

	select {
	case ch.(chan Prompt) <- Prompt{TeleCtx: ctx}:
	default:
		return false
	}

	return true
}
