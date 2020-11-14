package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"todoApi/api"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	api.StartServer(":8080")
}
