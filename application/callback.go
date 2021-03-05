package application

import (
	"github.com/projectxpolaris/youtrans/service"
	"github.com/sirupsen/logrus"
)

var youvideoCallbackLogger = logrus.WithField("scope", "YouVideoCallback")

var DefaultYouVideoCallback YouVideoCallBack = YouVideoCallBack{
	CompleteChan: make(chan *service.Task, 0),
}

type YouVideoCallBack struct {
	CompleteChan chan *service.Task
}

func (c *YouVideoCallBack) Start() {
	youvideoCallbackLogger.Info("Start youvideo callback watcher")
	for {
		select {
		case task := <-c.CompleteChan:
			youvideoCallbackLogger.WithField("task id", task.Id).Info("receive complete task")
			err := DefaultYouVideoClient.SendCompleteTask(task)
			if err != nil {
				youvideoCallbackLogger.Error(err)
			}
		}
	}
}
