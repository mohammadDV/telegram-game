package service

import (
	"github.com/mohammaddv/telegram-game/internal/repository"
)

type App struct {
	account *AccountService
}

func NewApp(account *AccountService) *App {
	return &App{account: account}
}