package file

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"tweet-eraser/helpers"
)

type TweetTime time.Time

const TweeterTimeFormat = time.RubyDate

type Tweets struct {
	Tweet `json:"tweet"`
}

type Tweet struct {
	Id        string    `json:"id"`
	CreatedAt TweetTime `json:"created_at"`
}

func (tt *TweetTime) UnmarshalJSON(b []byte) error {
	timeString := strings.Trim(string(b), `"`)
	t, err := time.Parse(TweeterTimeFormat, timeString)
	if err == nil {
		*tt = TweetTime(t)
		return nil
	}
	return errors.New(fmt.Sprintf("Invalid date format: %s", timeString))
}

func Decode(f string) ([]Tweet, error) {
	file, err := os.OpenFile(f, os.O_RDWR, 644)
	if err != nil {
		return nil, err
	}
	defer helpers.CloseFile(file)

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	_, err = decoder.Token()
	if err != nil {
		return nil, err
	}

	var tweets []Tweet
	for decoder.More() {
		elm := &Tweets{}
		err := decoder.Decode(elm)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, elm.Tweet)
	}

	return tweets, nil
}

func (tt *TweetTime) ToTime() time.Time {
	return time.Time(*tt)
}
