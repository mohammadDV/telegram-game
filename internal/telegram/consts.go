package telegram

import (
	"time"
	"github.com/mohammaddv/telegram-game/internal/entity"
	"gopkg.in/telebot.v3"
)

const (
	DefaultInputTimeout = time.Second * 30
	DefaultTimeoutText = "Timeout, please do it now"

	ConfirmText = "Confirm"
	DeclineText = "Decline"
)

func getAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}

var (
	selector = &telebot.ReplyMarkup{}
	brnEditDisplayName = selector.Data("Edit Display Name", "edit_display_name")
)