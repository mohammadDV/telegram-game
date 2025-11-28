package telegram

import (
	"time"
	"github.com/mohammaddv/telegram-game/internal/entity"
	"gopkg.in/telebot.v3"
)

const (
	DefaultInputTimeout = time.Second * 5
	DefaultTimeoutText = "Timeout, please do it now"
)

func getAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}