package main

import (
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	tgClient "telegramBot/clients/telegram"
	"telegramBot/consumer/event-consumer"
	"telegramBot/events/telegram"
	"telegramBot/lib/l"
	"telegramBot/storage/psql"
)

type JSONData struct {
	TgBotHost string `json:"tgBotHost"`
	ConnStr   string `json:"PSQLconnection"`
	BatchSize int    `json:"batchSize"`
}

const jsonFileName = "data.json"

// batchSize - updatesBatchLimit, between 1 - 100, defaults to 100

func main() {
	
	// logs
	l.Start()

	var launchData JSONData

	openJSONfiles(jsonFileName, &launchData)

	s, err := psql.New(launchData.ConnStr)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(launchData.TgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, launchData.BatchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped ", err)
	}
}

func openJSONfiles(fileName string, launchData *JSONData) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open JSON file: ", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read JSON file: ", err)
	}

	if err := json.Unmarshal(byteValue, launchData); err != nil {
		log.Fatal("Failed to parse JSON file: ", err)
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

