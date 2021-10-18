package cmd

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

type config struct {
	TweeterBackup      string
	AuthorizationToken string
	AuthToken          string
	CSRFToken          string
	DaysOffset         float64
}

var (
	cfg     config
	logPath = "report.log"
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.TweeterBackup, "file", "", "Your tweeter backup file path")
	rootCmd.PersistentFlags().StringVar(&cfg.AuthorizationToken, "authorization", "", "Your authorization code")
	rootCmd.PersistentFlags().StringVar(&cfg.AuthToken, "authToken", "", "Your authToken from cookie")
	rootCmd.PersistentFlags().StringVar(&cfg.CSRFToken, "csrfToken", "", "Your csrf token from cookie (ct0)")
	rootCmd.PersistentFlags().Float64Var(&cfg.DaysOffset, "offset", 0, "Up until when should I delete the tweets.")
}

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "run the application",
	Run:   run,
}

func Execute() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		// Can we log an error before we have our logger? :)
		log.Error().Err(err).Msg("there was an error creating a file for our log")
		os.Exit(1)
	}
	multi := zerolog.MultiLevelWriter(consoleWriter, file)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	// I'm not sure if it's a good idea, but didn't want to do it within the command!
	ctx := context.Background()
	ctx = context.WithValue(ctx, "logger", logger)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Error().Err(err).Msg("there was an error running the command")
		os.Exit(1)
	}
}
