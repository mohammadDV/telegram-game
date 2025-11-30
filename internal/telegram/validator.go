package telegram

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"gopkg.in/telebot.v3"
)

func chooseValidator(choices ...string) Validator {

	return Validator{
		validator: func(msg *telebot.Message) bool {
			return lo.Contains(choices, msg.Text)
		},
		OnInvalid: func(msg *telebot.Message) string {
			return fmt.Sprintf("Invalid choice, please choose from %s", strings.Join(choices, ", "))
		},
	}
}
