package main

import (
	"context"
	"godibot-atp/app"
	"godibot-atp/pkg/utils/logger"
	"log"
	"os"
)

var product = "BotVerse.AI"

func init() {
	log.Println("========   golang_discord  =========")
	log.Printf("app     : %s", product)
	log.Println("version : v1")
	log.Println("release : 2025-06-02")
	log.Println("========    initialization    =========")
}

func main() {
	logLevel := os.Getenv("LOG")
	token := os.Getenv("TOKEN")

	if len(logLevel) == 0 {
		logger.Level("fatal", "main", "logLevel is nil")
	}
	log.Printf("logLevel: %s", logLevel)
	logger.Load(logLevel)

	if len(token) == 0 {
		logger.Level("fatal", "main", "token is nil")
	}

	log.Println("============= start =============")
	ctx := context.Background()

	apps := setting(token)
	application := app.NewApp(*apps)
	if err := application.Start(ctx); err != nil {
		logger.Level("fatal", "main", "FAILED on Start: "+err.Error())
	}

	log.Println("=============  run  =============")
	if err := application.Run(ctx); err != nil {
		logger.Level("fatal", "main", "FAILED on Run: "+err.Error())
	}
}
