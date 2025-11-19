package integrationtest

import (
	"fmt"
	"testing"

	"os"

	"github.com/mohammaddv/telegram-game/pkg/testhelper"
)

func TestMain(m *testing.M) {

	if !testhelper.Integration() {
		return
	}

	pool := testhelper.StartDockerPool()
	resource := testhelper.StartDockerInstance(pool, "redis", "latest")
	fmt.Println(resource.GetPort("6379/tcp"))
	fmt.Println("redis is running")
	defer resource.Close()

	exitCode := m.Run()

	os.Exit(exitCode)
}
