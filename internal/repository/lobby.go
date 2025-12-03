package repository

import (
	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/redis/rueidis"
)

var _ LobbyRepository = &LobbyRedisRepository{}

type LobbyRedisRepository struct {
	*RedisCommonBehaviour[entity.Lobby]
}

func NewLobbyRedisRepository(client rueidis.Client) *LobbyRedisRepository {
	return &LobbyRedisRepository{
		RedisCommonBehaviour: NewRedisCommonBehaviour[entity.Lobby](client),
	}
}
