package main

import (
	"os"

	fluentd "github.com/joonix/log"
	log "github.com/sirupsen/logrus"

	_ "net/http/pprof"

	"github.com/dunglas/mercure/hub"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	switch os.Getenv("LOG_FORMAT") {
	case "JSON":
		log.SetFormatter(&log.JSONFormatter{})
		return
	case "FLUENTD":
		log.SetFormatter(fluentd.NewFormatter())
	}
}

func main() {
	hub, db, err := hub.NewHubFromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	hub.Start()
	hub.Serve()
}
