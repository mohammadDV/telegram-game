package repository

import "github.com/mohammaddv/telegram-game/internal/entity"


var _ CommonBehaviour[entity.Entity] = &RedisCommonBehaviour[entity.Entity]{}

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, ent entity.Entity) error
}

type RedisCommonBehaviour[T entity.Entity] struct {
	client rueidis.Client
}

func NewRedisCommonBehaviour[T entity.Entity](client rueidis.Client) RedisCommonBehaviour[T] {
	return &RedisCommonBehaviour[T]{client: client}
}

func (r *RedisCommonBehaviour[T]) Get(ctx context.Context, id entity.ID) (T, error) {
	var t T
	cmd := r.client.B().JsonGet().Key(id.String()).Path(".").Build()
	val, err := r.client.Do(ctx, cmd).ToString()
	if err != nil {

		if errors.is(err, rueidis.Nil) {
			return t, ErrorNotFound
		}

		logrus.WithError(err).WithField("id", id.String()).Errorln("failed to get entity from redis")
		return t, err
	}
	return jsonhelper.Decode[T]([]byte(val)), nil
}