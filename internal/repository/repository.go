package repository

import (
	"context"
	"errors"

	"github.com/mohammaddv/telegram-game/internal/entity"
)

var (
	ErrorNotFound = errors.New("not found")
)

//go:generate mockery --name=AccountRepository
type AccountRepository interface {
	CommonBehaviour[entity.Account]
}

//go:generate mockery --name=LobbyRepository
type LobbyRepository interface {
	CommonBehaviour[entity.Lobby]
}

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, ent entity.Entity) error
}
