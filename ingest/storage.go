package ingest

import (
	"github.com/adamdecaf/twitter"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

var (
	TweetsKafkaTopic = "tweets.v1"
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

	// Internal methods
	produder sarama.AsyncProducer
}

// `NewKafkaStorage`...
func NewKafkaStorage(brokers []string) Storage {
	config := sarama.NewConfig()
	config.Net.TLS.Enable = false
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return kafkaStorage{
		Brokers: brokers,
		produder: producer,
	}
}

// `Store` sends messages off to kafka. It assumes they're formatted properly after calling
// `Serialize()` on the passed in values.
func (k kafkaStorage) Store(batch []twitter.Tweet) error {
	for _, tweet := range batch {
		b := tweet.Serialize()
		if b == nil {
			// Skip any tweets that don't properly serialize
			continue
		}

		k.produder.Input() <- &sarama.ProducerMessage{
			Topic: TweetsKafkaTopic,
			Key: sarama.StringEncoder(tweet.Id),
			Value: sarama.ByteEncoder(*b),
		}
	}
	return nil
}
