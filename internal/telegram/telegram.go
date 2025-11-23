package telegram

import (
	"github.com/mohammaddv/telegram-game/internal/service"
	"gopkg.in/telebot.v3"
	"github.com/sirupsen/logrus"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot
}

func NewTelegram(app *service.App,apiKey string) (*Telegram, error) {
	pref := telebot.Settings{
		Token:  apiKey,
		Poller: &telebot.LongPoller{Timeout: 60 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Errorln("failed to create bot")
		return nil, err
	}

	t := &Telegram{
		bot: bot,
		App: app,
	}

	t.setupHandlers()

	return t, nil


}

func (t *Telegram) Start() {
	t.bot.Start()
}