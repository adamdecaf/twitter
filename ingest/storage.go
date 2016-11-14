package ingest

import (
	"github.com/adamdecaf/twitter"
)

// An layer for storing tweets. It expects a batch of tweets instead of single tweets
// at a time.
type Storage interface {
	Store([]twitter.Tweet) error
}

// `kafkaStorage` writes batches of tweets to kafka.
type kafkaStorage struct {
	Storage

	Brokers []string
}

// `NewKafkaStorage`...
func NewKafkaStorage(brokers []string) Storage {
	return kafkaStorage{
		Brokers: brokers,
	}
}

// `Store`...
func (k kafkaStorage) Store(batch []twitter.Tweet) error {
	return nil
}
