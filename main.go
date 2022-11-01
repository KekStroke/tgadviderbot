package main

import (
	"flag"
	"log"

	tgClient "github.com/kekstroke/tgadviserbot/clients/telegram"
	eventconsumer "github.com/kekstroke/tgadviserbot/consumer/event-consumer"
	"github.com/kekstroke/tgadviserbot/events/telegram"
	"github.com/kekstroke/tgadviserbot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

// 5720163012:AAH682wfTFm2LhPFYIofJ0_MJ32VgYmyxNo
func main() {
	token := mustToken()

	tgClient := tgClient.New(tgBotHost, token)

	eventsProcessor := telegram.New(tgClient, files.New(storagePath))

	log.Println("service started...")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	// bot -tg-bot-token 'my token'
	token := flag.String("tg-bot-token", "", "telegram bot token access")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
