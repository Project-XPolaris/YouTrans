package service

import (
	"github.com/allentom/transcoder"
	"github.com/allentom/transcoder/ffmpeg"
	"github.com/projectxpolaris/youtrans/config"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
)

var Logger = log.New().WithFields(log.Fields{
	"scope": "TaskManager",
})
var DefaultTaskPool = TaskPool{
	Tasks: []*Task{},
}

type TaskPool struct {
	Tasks []*Task
	sync.RWMutex
}

func (p *TaskPool) CreatTask(option *TaskOption) error {
	id := "123"
	task := &Task{
		Id:       id,
		Option:   option,
		DoneChan: make(chan struct{}),
		Logger: log.New().WithFields(log.Fields{
			"scope": "Task",
			"id":    id,
		}),
	}
	go task.Run()
	p.Lock()
	p.Tasks = append(p.Tasks, task)
	p.Unlock()
	return nil
}

func (p *TaskPool) GetTaskById(id string) *Task {
	p.Lock()
	defer p.Unlock()
	for _, task := range p.Tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}
func (p *TaskPool) StopTask(id string) {
	task := p.GetTaskById(id)
	if task != nil && task.Status == "Running" {
		task.InterruptFlag = true
		task.Transcoder.Stop()
	}
	return
}

type Task struct {
	Id            string
	Option        *TaskOption
	Progress      transcoder.Progress
	Status        string
	StopFlag      bool
	DoneChan      chan struct{}
	Logger        *log.Entry
	InterruptFlag bool
	Transcoder    transcoder.Transcoder
}
type TaskOption struct {
	Option     ffmpeg.Options
	Overwrite  bool
	Format     string
	InputPath  string
	OutputPath string
}
type DiscardCloser struct {
	io.Writer
}

func (DiscardCloser) Close() error {
	return nil
}
func (t *Task) Run() error {
	opts := ffmpeg.Options{
		OutputFormat: &t.Option.Format,
		Overwrite:    &t.Option.Overwrite,
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   config.DefaultConfig.FfmpegBin,
		FfprobeBinPath:  config.DefaultConfig.FfprobeBin,
		ProgressEnabled: true,
	}
	trans := ffmpeg.
		New(ffmpegConf).
		Input(t.Option.InputPath).
		Output(t.Option.OutputPath).
		WithOptions(opts)
	prg, err := trans.Start(opts)
	t.Transcoder = trans
	if err != nil {
		return err
	}

	t.Status = "Running"
	go func() {
		for {
			msg, isClose := <-prg
			if isClose {
				t.Progress = msg
				continue
			}
			if t.InterruptFlag {
				t.Logger.Info("interrupt transcode")
				t.Status = "Stop"
			} else {
				t.Logger.Info("transcode done")
				t.Status = "Done"
			}
			break
		}
		t.DoneChan <- struct{}{}
	}()

	return nil
}