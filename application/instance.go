package application

import (
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/youtrans/config"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var Logger = log.New().WithField("scope", "Handler")

func RunApplication() {
	e := haruka.NewEngine()
	e.UseCors(cors.AllowAll())
	e.Router.POST("/tasks", createTransHandler)
	e.Router.GET("/tasks", taskListHandler)
	e.Router.POST("/tasks/stop", stopTransHandler)
	e.Router.GET("/ffmpeg/codec", codecListHandler)
	e.RunAndListen(config.DefaultConfig.Addr)
}
