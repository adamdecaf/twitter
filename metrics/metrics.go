package metrics

import (
	"github.com/armon/go-metrics"
	"time"
)

var (
	sink = metrics.NewInmemSink(10*time.Second, time.Minute)
	sig = metrics.DefaultInmemSignal(sink)
)

func init() {
	metrics.NewGlobal(metrics.DefaultConfig("ingest"), sink)
}

// Define a counter
type Counter struct {
	Name string
}
func (c Counter) Add(v float32) {
	sink.IncrCounter([]string{}, v)
}
func (c Counter) AddI(v int) {
	c.Add(float32(v))
}

// Create inmem counters
func NewCounter(name string) *Counter {
	return &Counter{
		Name: name,
	}
}
