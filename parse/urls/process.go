package urls

import (
	"github.com/adamdecaf/twitter"
	"github.com/adamdecaf/twitter/metrics"
	"github.com/mvdan/xurls"
	"log"
)

var (
	UrlsCounter = metrics.NewCounter("tweets.urls-found")
)

// `process` grabs the raw bytes from a message off kafka and
// finds urls within the tweet body.
// `b` is json of twitter.Tweet
func process(b []byte) {
	tweet, err := twitter.ReadTweet(b)
	if err != nil {
		log.Printf("unable to read tweet, err=%v", err)
	}

	urls := findUrls(tweet.Text)
	UrlsCounter.AddI(len(urls))
}

// `findUrls` finds urls
func findUrls(s string) []string {
	return xurls.Relaxed.FindAllString(s, -1)
}
