package cmd

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the game",
	Run: serve,
}

func serve(_ *cobra.Command, _ []string) {
	_ = godotenv.Load();
	fmt.Println("serve called")
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
