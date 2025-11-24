package repository

import (
	"context"
	"errors"

	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/mohammaddv/telegram-game/pkg/jsonhelper"
	"github.com/redis/rueidis"
	"github.com/sirupsen/logrus"
)

var _ CommonBehaviour[entity.Entity] = &RedisCommonBehaviour[entity.Entity]{}

type RedisCommonBehaviour[T entity.Entity] struct {
	client rueidis.Client
}

func NewRedisCommonBehaviour[T entity.Entity](client rueidis.Client) *RedisCommonBehaviour[T] {
	return &RedisCommonBehaviour[T]{client: client}
}

func (r *RedisCommonBehaviour[T]) Get(ctx context.Context, id entity.ID) (T, error) {
	var t T
	cmd := r.client.B().JsonGet().Key(id.String()).Path(".").Build()
	val, err := r.client.Do(ctx, cmd).ToString()
	if err != nil {

		if errors.Is(err, rueidis.Nil) {
			return t, ErrorNotFound
		}

		logrus.WithError(err).WithField("id", id.String()).Errorln("failed to get entity from redis")
		return t, err
	}
	return jsonhelper.Decode[T]([]byte(val)), nil
}

func (r *RedisCommonBehaviour[T]) Save(ctx context.Context, ent entity.Entity) error {
	cmd := r.client.B().JsonSet().Key(ent.EntityID().String()).Path("$").Value(string(jsonhelper.Encode(ent))).Build()

	if err := r.client.Do(ctx, cmd).Error(); err != nil {
		logrus.WithError(err).WithField("ent", ent).Errorln("failed to save entity to redis")
		return err
	}
	return nil
}
