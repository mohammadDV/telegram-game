package cmd

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/mohammaddv/telegram-game/internal/repository/redis"
	"github.com/sirupsen/logrus"
	"os"
	"github.com/mohammaddv/telegram-game/internal/repository"
	"github.com/mohammaddv/telegram-game/internal/service"
	"github.com/mohammaddv/telegram-game/internal/telegram"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the game",
	Run: serve,
}

func serve(_ *cobra.Command, _ []string) {
	_ = godotenv.Load();

	// setup repository
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		logrus.WithError(err).Errorln("failed to create redis client")
		return
	}

	// setup service
	accountRepository := repository.NewAccountRedisRepository(redisClient)
	App := service.NewApp(
		service.NewAccountService(accountRepository),
	)

	// setup telegram
	telegram, err := telegram.NewTelegram(App, os.Getenv("BOT_API"))
	if err != nil {
		logrus.WithError(err).Errorln("failed to create telegram")
		return
	}

	// start telegram
	telegram.Start()

	fmt.Println("serve called")
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
