package ingest

import (
	"errors"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"time"
)

// `clientConfig` is the set of configuration options for a `Client`. These are inline with
// Twitter's best practices.
// More Information: https://dev.twitter.com/streaming/overview/connecting
type clientConfig struct {
	Rate time.Duration
	BufferSize int64
}

// `NewClientConfig` creates a new `clientConfig` with the default options.
func NewConfig() clientConfig {
	return clientConfig{
		Rate: 0 * time.Second,
		BufferSize: 1000,
	}
}

// `Client` is a channel that streams events from twitter out for consumers to access.
// This is directly reading off the twitter api so it's important to not block this.
type Client struct {
	Config clientConfig

	// The raw events coming off twitter's api
	Events chan interface{}
}

// `NewClient` returns a new `Client` instance each time it's called. It will return a nil
// client and an error if something goes wrong. The errors are not meant to be machine readable.
func NewClient(creds credentials, config clientConfig) (*Client, error) {
	anaconda.SetConsumerKey(creds.consumerKey)
	anaconda.SetConsumerSecret(creds.consumerSecret)

	api := anaconda.NewTwitterApi(creds.accessToken, creds.accessSecret)
	api.EnableThrottling(config.Rate, config.BufferSize)

	params := url.Values{}
	params.Add("language", "en")

	stream := api.PublicStreamSample(params)
	if stream == nil {
		return nil, errors.New("Error getting twitter stream")
	}

	return &Client{
		Config: config,
		Events: stream.C,
	}, nil
}
