package integrationtest

import (
	"fmt"
	"testing"

	"os"

	"github.com/mohammaddv/telegram-game/internal/repository/redis"
	"github.com/mohammaddv/telegram-game/pkg/testhelper"
	"github.com/ory/dockertest/v3"
)

var redisPort string

func TestMain(m *testing.M) {

	if !testhelper.IsIntegration() {
		return
	}

	pool := testhelper.StartDockerPool()
	resource := testhelper.StartDockerInstance(pool, "redis", "latest",
		func(res *dockertest.Resource) error {
			port := res.GetPort("6379/tcp")
			_, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", port))
			return err
		},
	)
	redisPort = resource.GetPort("6379/tcp")
	fmt.Println("redis is running")
	defer resource.Close()

	exitCode := m.Run()

	os.Exit(exitCode)
}
