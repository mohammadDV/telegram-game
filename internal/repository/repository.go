package repository

import (
	"errors"
	"github.com/mohammaddv/telegram-game/internal/entity"
	"context"
)


 var (
	ErrorNotFound = errors.New("not found")
 )

 type AccountRepository interface {
	CommonBehaviour[entity.Account]
}

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, ent entity.Entity) error
}