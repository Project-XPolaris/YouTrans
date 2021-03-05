package main

import (
	"github.com/projectxpolaris/youtrans/application"
	"github.com/projectxpolaris/youtrans/config"
	log "github.com/sirupsen/logrus"
)

var Logger = log.WithField("scope", "Main")

func main() {
	err := config.LoadConfig()
	if err != nil {
		Logger.WithField("action", "load config").Fatal(err)
	}
	Logger.WithField("action", "load config").Info("success load config")
	go application.DefaultYouVideoCallback.Start()
	Logger.Info("YouVideo Url = " + config.DefaultConfig.YouVideoUrl)
	application.RunApplication()
}
