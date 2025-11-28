package telegram

import (
	"context"

	"time"
	"errors"
	"fmt"
	"github.com/google/uuid"
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
		c.Set("is_just_created", accountSaved)

		return next(c)
	}
}


func (t *Telegram) OnError(err error, c telebot.Context) {
	if errors.Is(err, ErrorInputTimeout) {
		return
	}

	errorId := uuid.New().String()

	logrus.WithError(err).WithField("error_id", errorId).Errorln("telegram error")

	c.Reply(fmt.Sprintf("An error occurred. Please try again later. Error ID: %s", errorId))
}

