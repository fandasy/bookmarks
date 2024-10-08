package eventconsumer

import (
	"log"
	"telegramBot/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {

	log.Print("Consumer started")

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

		if err := c.hendleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

func (c *Consumer) hendleEvents(events []events.Event) error {

	for _, event := range events {
		go func() {
			log.Printf("got new event: %s", event.Text)

			if err := c.processor.Process(event); err != nil {
				log.Printf("can't handle event: %s", err.Error())
			}
		}()
	}

	return nil
}
