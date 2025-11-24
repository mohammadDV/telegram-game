package service

import (
	"context"

	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/mohammaddv/telegram-game/internal/repository"
	"errors"
	"time"
)

type AccountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

var (
	AccountStateNew = "home"
	AccountStateActive = "active"
	AccountStateInactive = "inactive"
)

func (a *AccountService) UpdateOrCreate(ctx context.Context, account entity.Account) (entity.Account, bool, error) {


	savedAccount, err := a.accountRepository.Get(ctx, account.EntityID())
	if err == nil {
		if account.Username != savedAccount.Username || account.FirstName != savedAccount.FirstName {
			savedAccount.Username = account.Username
			savedAccount.FirstName = account.FirstName
			return savedAccount, false, a.accountRepository.Save(ctx, savedAccount)
		}
		return savedAccount, false, nil
	}

	if errors.Is(err, repository.ErrorNotFound) {
		account.State = AccountStateNew
		account.JoinedAt = time.Now()
		return account, true, a.accountRepository.Save(ctx, account)
	}

	return entity.Account{}, false, err
}