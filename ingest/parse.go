package ingest

import (
	"github.com/adamdecaf/twitter"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"time"
)

func parse(tweet anaconda.Tweet) (*twitter.Tweet, *twitter.User) {
	t := parseTweetDetails(tweet)
	u := parseUserDetails(tweet)
	return t, u
}

// This won't return a valid User record
func parseTweetDetails(tweet anaconda.Tweet) *twitter.Tweet {
	when, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.CreatedAt)
	if err != nil {
		log.Printf("failed tweet time parse: when=%s | err=%s\n", when, err)
		return nil
	}

	return &twitter.Tweet{
		Id: tweet.IdStr,
		UserId: tweet.User.IdStr,
		Text: tweet.Text,
		CreatedAt: when,
	}
}

// This just reads out user information
func parseUserDetails(tweet anaconda.Tweet) *twitter.User {
	when, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.User.CreatedAt)
	if err != nil {
		log.Printf("failed user time parse: when=%s | err=%s\n", when, err)
		return nil
	}

	return &twitter.User{
		Id: tweet.User.IdStr,
		Name: tweet.User.Name,
		ScreenName: tweet.User.ScreenName,
		CreatedAt: when,
	}
}
