package integrationtest

import (
	"testing"

	"github.com/mohammaddv/telegram-game/internal/repository/redis"
	"github.com/mohammaddv/telegram-game/pkg/testhelper"
	"github.com/stretchr/testify/assert"
	"context"
	"github.com/mohammaddv/telegram-game/internal/machmaking"
	"fmt"

)

func TestMatchmaking(t *testing.T) {
	// if !testhelper.IsIntegration() {
	// 	t.Skip("integration test")
	// }
	ctx := context.Background()
	timeout := time.Second * 10

	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)


	mm := machmaking.NewRedisMatchmaking(redisClient, repository.NewLobbyRedisRepository(redisClient)) 
		
	var wg sync.WaitGroup
	testJoin := func(id int64) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lobby, err := mm.Join(ctx, id, timeout)
			assert.NoError(t, err)
			assert.NotEqual(t, "", lobby)

			wg.Done()
		}()

	}

	testJoin(13)
	testJoin(14)
	testJoin(15)
	testJoin(16)

	<-time.After(500 * time.Millisecond)

	assert.Equal(t, int64(4), zCount(ctx, redisClient, "matchmaking"))

	lobby, err := mm.Join(ctx, 17, timeout)
	assert.NoError(t, err)
	assert.NotEqual(t, "", lobby.ID)
	wg.Wait()

}

func TestMatchmakingWithManyLobbies(t *testing.T) {
	// if !testhelper.IsIntegration() {
	// 	t.Skip("integration test")
	// }
	ctx := context.Background()
	timeout := time.Second * 10

	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)


	mm := machmaking.NewRedisMatchmaking(redisClient, repository.NewLobbyRedisRepository(redisClient)) 
		
	counter := newLobbyCounter()
	var wg sync.WaitGroup
	testJoin := func(id int64) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lobby, err := mm.Join(ctx, id, timeout)
			assert.NoError(t, err)
			assert.NotEqual(t, "", lobby)
			counter.Incr(lobby.ID)
			wg.Done()
		}()

	}

	for i := 0; i < 1000; i++ {
		testJoin(int64(i))
	}

	wg.Wait()

	assert.Equal(t, counter.counter, 200)
	for _, count := range counter.counter {
		assert.Equal(t, 5, count)
	}

}

func zCount(ctx context.Context, redisClient rueidis.Client, key string) int64 {
	count, err := redisClient.Do(context.Background(), 
		redisClient.B().Zcount().Keys("matchmaking").Min("-inf").Max("+inf").Build(),
	).ToInt64()

	assert.NoError(t, err)
	return count
}

func keys(ctx context.Context, redisClient rueidis.Client, pattern string) []string {
	keys, err := redisClient.Do(context.Background(), 
		redisClient.B().Keys().Pattern(pattern).Build()
	).AsStrSlice()

	assert.NoError(t, err)
	return keys
}