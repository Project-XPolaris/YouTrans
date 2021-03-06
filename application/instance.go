package application

import (
	"github.com/allentom/haruka"
	"github.com/allentom/haruka/middleware"
	"github.com/projectxpolaris/youtrans/config"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var Logger = log.New().WithField("scope", "Handler")

func RunApplication() {
	e := haruka.NewEngine()
	e.UseCors(cors.AllowAll())
	e.UseMiddleware(middleware.NewLoggerMiddleware())
	e.Router.GET("/tasks", taskListHandler)
	e.Router.POST("/tasks", createTransHandler)
	e.Router.POST("/tasks/stop", stopTransHandler)
	e.Router.GET("/ffmpeg/codec", getCodecsHandler)
	e.Router.GET("/ffmpeg/formats", getFormatsHandler)
	e.Router.GET("/info", infoHandler)
	e.RunAndListen(config.Instance.Addr)
}
