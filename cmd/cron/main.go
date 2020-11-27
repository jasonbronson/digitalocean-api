package main

import (
	"docontroller/jobs"
	"docontroller/utils"
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"

	"github.com/roylee0704/gron"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		utils.LogError(errors.New("Cannot find .env file skipping "))
	}

	cron := gron.New()

	// Job for IAP Receipt checker
	if d := os.Getenv("CRONJOB_CLEANUP"); d != "" {
		utils.Log("cronjob cleanup")
		duration, err := time.ParseDuration(d)
		if err != nil {
			utils.LogError(err)
		}

		cron.Add(gron.Every(duration), &jobs.Cleanup{})
	}

	cron.Start()
	log.Println("=====cron system started======")

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
}
