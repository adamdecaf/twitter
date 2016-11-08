package ingest

import (
	"errors"
	"os"
)

var (
	// Credentials
	CONSUMER_KEY	= os.Getenv("CONSUMER_KEY")
	CONSUMER_SECRET = os.Getenv("CONSUMER_SECRET")
	ACCESS_TOKEN	= os.Getenv("ACCESS_TOKEN")
	ACCESS_SECRET	= os.Getenv("ACCESS_SECRET")
)

// `Credentials` are the set of credentials required for accessing the twitter api.
// More information:
type credentials struct {
	consumerKey string
	consumerSecret string

	accessToken string
	accessSecret string
}

func NewCredentials() (*credentials, error) {
	// Check credentials
	if CONSUMER_KEY == "" {
		err := errors.New("missing consumer keys")
		return nil, err
	}
	if CONSUMER_SECRET == "" {
		err := errors.New("missing consumer secret")
		return nil, err
	}
	if ACCESS_TOKEN == "" {
		err := errors.New("missing access token")
		return nil, err
	}
	if ACCESS_SECRET == "" {
		err := errors.New("missing access secret")
		return nil, err
	}

	return &credentials{
		consumerKey: CONSUMER_KEY,
		consumerSecret: CONSUMER_SECRET,
		accessToken: ACCESS_TOKEN,
		accessSecret: ACCESS_SECRET,
	}, nil
}
