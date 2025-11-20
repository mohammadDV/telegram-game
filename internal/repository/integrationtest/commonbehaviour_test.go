package integrationtest


import (
	"fmt"
	"testing"

	"github.com/mohammaddv/telegram-game/internal/repository/redis"
	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/mohammaddv/telegram-game/internal/repository"
	"context"
)

type testType struct {
	ID string
	Name string
}

func (t *testType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviourSetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)

	cb := repository.NewRedisCommonBehaviour[testType](redisClient)
	err = cb.Save(context.Background(), &testType{ID: "14", Name: "hassan"})
	assert.NoError(t, err)

	val, err := cb.Get(context.Background(), entity.NewID("testType", "14"))
	assert.NoError(t, err)
	assert.Equal(t, val.Name, "hassan")
	assert.Equal(t, val.ID, "14")
	
	
}