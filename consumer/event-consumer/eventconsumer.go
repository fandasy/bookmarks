package eventconsumer

import (
	"log"
	"runtime"
	"sync"
	"telegramBot/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

var jobs = make(chan []events.Event)

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {

	defer close(jobs)

	log.Print("Consumer started")

	c.generateGoroutines()

	wg := sync.WaitGroup{}

	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			jobs <- gotEvents
		}()

		wg.Wait()
	}
}

func (c *Consumer) generateGoroutines() {

	numGoroutines := runtime.NumCPU()

	for i := 1; i <= numGoroutines; i++ {
		go func() {
			for {
				select {
				case gotEvents := <-jobs:
					if err := c.hendleEvents(gotEvents); err != nil {
						log.Print(err)

						continue
					}
					time.Sleep(1 * time.Second)
				default:
					continue
				}
			}
		}()
	}
}

func (c *Consumer) hendleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
