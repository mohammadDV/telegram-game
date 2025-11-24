package integrationtest

import (
	"fmt"
	"testing"

	"context"

	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/mohammaddv/telegram-game/internal/repository"
	"github.com/mohammaddv/telegram-game/internal/repository/redis"
	"github.com/stretchr/testify/assert"
)

type testType struct {
	ID   string
	Name string
}

func (t testType) EntityID() entity.ID {
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
	fmt.Println(val)

	val, err = cb.Get(context.Background(), entity.NewID("testType", "16"))
	assert.ErrorIs(t, err, repository.ErrorNotFound)

	// Print key info for debugging
	fmt.Printf("\nTest created key: %s\n", entity.NewID("testType", "14").String())
	fmt.Printf("To view in Redis: redis-cli -h localhost -p %s GET %s\n", redisPort, entity.NewID("testType", "14").String())
	fmt.Printf("To view all keys: redis-cli -h localhost -p %s KEYS '*'\n", redisPort)

}
