package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"tgrepost/bot"
)

func main() {
	config, err := bot.NewConfig()
	if err != nil {
		log.Fatalf("creating config: %v", err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := bot.RunBot(ctx, config); err != nil {
		log.Fatalf("running bot: %v", err)
	}
}
