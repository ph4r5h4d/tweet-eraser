package file

import (
	"bufio"
	"os"
	"tweet-eraser/helpers"
)

var (
	tweetJSON = "./.tmp/tweets.json"
)

// FixFile rewrite the original tweet,js file into a correct Json file. Not the best way, I know, but works for now.
func FixFile(input string) error {
	file, err := os.OpenFile(input, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer helpers.CloseFile(file)

	tweets, _ := os.Create(tweetJSON)
	defer helpers.CloseFile(tweets)

	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
		if firstLine {
			// skip the first line and replace it with "[" to convert the js file into a correct json file.
			scanner.Text()
			_, err := tweets.WriteString("[\n")
			if err != nil {
				return err
			}
			firstLine = false
			continue
		}

		_, err := tweets.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
