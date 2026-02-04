package main

import "shared/logger"

func main() {
	log := logger.New()
	defer log.Sync()

	log.Info("user service started")
}
