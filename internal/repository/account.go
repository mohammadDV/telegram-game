package repository

import (
	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/redis/rueidis"
)


type AccountRedisRepository struct {
	*RedisCommonBehaviour[entity.Account]
}

func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
	return &AccountRedisRepository{
		RedisCommonBehaviour: NewRedisCommonBehaviour[entity.Account](client),
	}
}