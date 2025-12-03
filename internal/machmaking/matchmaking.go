package machmaking

import (
	"context"
	_ "embed"
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/mohammaddv/telegram-game/internal/repository"
	"github.com/redis/rueidis"
	"github.com/sirupsen/logrus"
	"fmt"
)

//go:embed matchmaking.lua
var matchmakingScript string

var (
	ErrBadRedisResponse = errors.New("bad redis response")
	ErrTimeout = errors.New("timeout")
)

type Matchmaking interface {
	Join(ctx context.Context, userId int64, timeout time.Duration) (entity.Lobby, error)
	Leave(ctx context.Context, lobby entity.Lobby) error
}

// var _ Matchmaking = &RedisMatchmaking{}

type RedisMatchmaking struct {
	redisClient       rueidis.Client
	matchmakingScript *rueidis.Lua
	lobby repository.LobbyRepository
}

func NewRedisMatchmaking(redisClient rueidis.Client, lobby repository.LobbyRepository) *RedisMatchmaking {
	script := rueidis.NewLuaScript(matchmakingScript)
	return &RedisMatchmaking{
		redisClient:       redisClient,
		matchmakingScript: script,
		lobby: lobby,
	}
}

type joinLobbyPubSubResponse struct {
	err error
	lobbyId string
}

func (r *RedisMatchmaking) Join(ctx context.Context, userId int64, timeout time.Duration) (entity.Lobby, error) {
	waitingLobbyCtx, lobbyContextCancel := context.WithTimeout(context.Background(), timeout)

	defer lobbyContextCancel()
	responseChan := make(chan joinLobbyPubSubResponse, 1)

	go r.redisClient.Receive(waitingLobbyCtx, r.redisClient.B().Subscribe().Channel("matchmaking").Build(), func(msg rueidis.PubSubMessage) {
		split := strings.Split(msg.Message, ":")
		lobbyId := split[0]
		users := lo.Map(strings.Split(split[1], ","), func(item string, _ int) int64 {
			id, _ := strconv.ParseInt(item, 10, 64)
			return id
		})

		if !slices.Contains(users, userId) {
			return
		}

		responseChan <- joinLobbyPubSubResponse{
			lobbyId: lobbyId,
		}

		
	})
	
	resp, err := r.matchmakingScript.Exec(ctx, r.redisClient,
		[]string{"matchmaking", "matchmaking"},
		[]string{"4", strconv.FormatInt(
			time.Now().Add(-time.Minute*10).Unix(), 10),
			uuid.New().String(),
			strconv.FormatInt(userId, 10),
			strconv.FormatInt(time.Now().Unix(), 10),
			strconv.FormatInt(userId, 10),
			strconv.FormatInt(time.Now().Unix(), 10),
		}).ToArray()
	if err != nil {
		logrus.WithError(err).Errorln("failed to join matchmaking")
		return entity.Lobby{}, err
	}

	// inside a queue: we must listen to the pubsub channel for the lobby id
	if len(resp) == 1 {

		select {
		case pubSubResponse := <-responseChan:
			return r.lobby.Get(ctx, entity.NewID("lobby", pubSubResponse.lobbyId))
		case <-waitingLobbyCtx.Done():
			return entity.Lobby{}, ErrTimeout
		}
	}

	// found a lobby: we have created a lobby and returned the lobby id and the participants
	if len(resp) == 3 {
		lobbyId, _ := resp[1].ToString()
		return r.lobby.Get(ctx, entity.NewID("lobby", lobbyId))
	}

	return entity.Lobby{}, ErrBadRedisResponse
}
