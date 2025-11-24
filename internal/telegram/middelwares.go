package telegram

import (
	"context"

	"time"

	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func (t *Telegram) registerMiddelware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:        c.Sender().ID,
			Username:  c.Sender().Username,
			FirstName: c.Sender().FirstName,
			JoinedAt:  time.Now(),
		}

		account, accountSaved, err := t.App.Account().UpdateOrCreate(context.Background(), acc)
		if err != nil {
			logrus.WithError(err).WithField("account", acc).Errorln("failed to update or create account")
			return err
		}

		c.Set("account", account)
		c.Set("accountSaved", accountSaved)

		return next(c)
	}
}
