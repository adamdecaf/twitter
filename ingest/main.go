package ingest

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
	"os"
)

func main() {
	config := NewConfig()
	credentials, err := NewCredentials()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if credentials == nil {
		log.Println("No credentials found")
		os.Exit(1)
	}

	client, err := NewClient(*credentials, config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Read off tweets forever, die if something panics.
	// todo: leaks underlying twitter library
	for {
		item := <- client.Events
		tweet, ok := item.(anaconda.Tweet)
		if ok {
			go parseAndStore(tweet)
		}
	}
}

func parseAndStore(tweet anaconda.Tweet) {
	t, _ := parse(tweet)

	// store tweet
	if t != nil {
		storeTweet(*t)
	}
}
