package ingest

import (
	"github.com/adamdecaf/twitter"
	"github.com/adamdecaf/twitter/metrics"
	"log"
	"os"
	"strings"
)

// todo
// Hook into metrics registry
// Config.MetricRegistry = (metrics.Registry)
// From: "github.com/rcrowley/go-metrics"

var (
	DefaultStorageBatchSize = 100
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

	brokers := strings.Split(os.Getenv("BROKERS_LIST"), ",")
	storage := NewKafkaStorage(brokers)
	batch := make([]twitter.Tweet, 0, DefaultStorageBatchSize)
	counter := metrics.NewCounter("incoming-tweets")

	// Read off tweets forever, die if something panics.
	for {
		item := <- client.Events
		t, _ := parse(item)

		// Did we get a tweet?
		if t != nil {
			batch = append(batch, *t)
		}

		// Store tweets once we've got a batch inmem
		if len(batch) >= DefaultStorageBatchSize {
			counter.AddI(len(batch))
			err := storage.Store(batch)
			if err != nil {
				log.Println(err)
			}
			batch = nil
		}
	}
}
