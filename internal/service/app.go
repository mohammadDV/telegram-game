package service

type App struct {
	account *AccountService
}

func NewApp(account *AccountService) *App {
	return &App{account: account}
}

func (a *App) Account() *AccountService {
	return a.account
}
