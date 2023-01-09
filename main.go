package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/conductor-sdk/go-sdk-examples/internal/worker"
)

func setupLogSettings() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.Info("Finished setting up log settings")
}

func main() {
	setupLogSettings()
	worker.StartWorkers()
	time.Sleep(5 * time.Second)
}
