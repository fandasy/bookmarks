package main

import (
	"context"
	"flag"
	"log"

	tgClient "telegramBot/clients/telegram"
	"telegramBot/consumer/event-consumer"
	"telegramBot/events/telegram"
	"telegramBot/storage/psql"
)

// "telegramBot/storage/files"

const (
	tgBotHost = "api.telegram.org"
	connStr   = "user=postgres dbname=tgBot password=iv83 host=localhost port=5432 sslmode=disable"
	batchSize = 100
)

func main() {
	// s := files.New(storagePath)
	s, err := psql.New(connStr)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
