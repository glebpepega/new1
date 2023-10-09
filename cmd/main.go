package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/glebpepega/new1/internal/apiserver"
	"github.com/glebpepega/new1/internal/db"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	dbInstance, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(dbInstance, logger)

	go s.ConfigureServer()

	sigintChan := make(chan os.Signal, 1)
	signal.Notify(sigintChan, os.Interrupt)

	<-sigintChan
	close(sigintChan)

	s.Logger.Info("graceful shutdown")

	if err := s.Fiber.ShutdownWithTimeout(time.Second * 10); err != nil {
		s.Logger.Fatal(err)
	}
}
