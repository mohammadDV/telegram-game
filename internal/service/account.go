package service

import (
	"context"

	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/mohammaddv/telegram-game/internal/repository"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	accountRepository repository.AccountRepository
}

func (a *AccountService) UpdateOrCreate(ctx context.Context, account entity.Account) error {
	err := a.accountRepository.Save(ctx, account)
	if err != nil {
		logrus.WithError(err).WithField("account", account).Errorln("failed to update or create account")
		return err
	}
	return nil
}