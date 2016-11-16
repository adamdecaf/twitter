package urls

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"
)

// todo
// Hook into metrics registry
// Config.MetricRegistry = (metrics.Registry)
// From: "github.com/rcrowley/go-metrics"

var (
	TweetsKafkaTopic = "tweets.v1"
)

func main() {
	brokers := strings.Split(os.Getenv("BROKERS_LIST"), ",")

	config := sarama.NewConfig()
	config.Net.TLS.Enable = false
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("unable to connect to kafka, err=%s", err)
	}

	// try to close cleanly when we die
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(TweetsKafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("error making consumer, err=%s", err)
	}

	// try to close gracefully
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		msg := <-partitionConsumer.Messages()
		go process(msg.Value)
	}

}
