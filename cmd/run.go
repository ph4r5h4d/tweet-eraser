package cmd

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"tweet-eraser/eraser"
	"tweet-eraser/file"
	"tweet-eraser/helpers"
)

var (
	tweetJsPath       = ".tmp/twdata/data/tweets.js"
	zipExtractionPath = ".tmp/twdata"
	tweetJsonPath     = "./.tmp/tweets.json"
)

func run(cmd *cobra.Command, args []string) {
	logger := cmd.Context().Value("logger").(zerolog.Logger)

	//unzip the package
	logger.Info().Msg("Starting to unzip the package...")
	err := file.Unzip(cfg.TweeterBackup, zipExtractionPath)
	if err != nil {
		logger.Error().Err(err).Msg("failed to unzip the package")
		return
	}
	logger.Info().Msg("Packaged extracted successfully")

	logger.Info().Msg("Let me find the tweets file and apply a small fix...")
	err = file.FixFile(tweetJsPath)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fix the tweet.js file")
		return
	}
	logger.Info().Msg("Alright, done.")

	logger.Info().Msg("Let's do some cleanup...")
	err = file.RemoveDir(zipExtractionPath)
	if err != nil {
		logger.Error().Err(err).Msg("cleanup failed by why?")
		return
	}
	logger.Info().Msg("Ah cleanup is done.")

	tweets, err := file.Decode(tweetJsonPath)
	if err != nil {
		logger.Error().Err(err).Msg("cleanup failed by why?")
		return
	}

	td := eraser.NewTweeterData()
	td.AuthToken(cfg.AuthToken)
	td.CSRFToken(cfg.CSRFToken)
	td.AuthorizationToken(cfg.AuthorizationToken)

	client := &http.Client{}
	rand.Seed(time.Now().UnixNano())
	for _, tweet := range tweets {
		if helpers.IsDeletable(tweet.CreatedAt.ToTime(), cfg.DaysOffset) {

			fmt.Println()
			logger.Info().Msg("Trying to delete id: " + tweet.Id + " Created at: " + tweet.CreatedAt.ToTime().String())
			status, _, err := td.DeleteTweet(tweet.Id, client)

			if err != nil {
				logger.Error().Msg("Error while deleting id: " + tweet.Id)
				logger.Error().Msg(">>>" + err.Error())
			}

			switch status {
			case 200:
				logger.Info().Msg("Id: " + tweet.Id + " deleted")
			case 404:
				logger.Warn().Msg("Id: " + tweet.Id + " doesn't exists or already deleted")
			case 403:
				logger.Warn().Msg("we hit a tweet we can't delete for some reason. continuing...")
			default:
				logger.Debug().Msg("Unacceptable status code: " + strconv.Itoa(status))
			}

			r := rand.Intn(2) + 1
			time.Sleep(time.Duration(r) * time.Second)
		}
	}

	logger.Info().Msg("Well I think we are done. Let's do one last cleanup...")
	err = file.RemoveFile(tweetJsonPath)
	if err != nil {
		logger.Error().Err(err).Msg("cleanup failed by why?")
	}
	logger.Info().Msg("I'm done, happy cleaning day :))")
}
