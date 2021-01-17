package application

import "github.com/projectxpolaris/youtrans/service"

type TaskTemplate struct {
	Id      string  `json:"id"`
	Process float64 `json:"process"`
	Input   string  `json:"input"`
	Output  string  `json:"output"`
	Status  string  `json:"status"`
}

func (t *TaskTemplate) Assign(task *service.Task) {
	t.Id = task.Id
	t.Output = task.Option.OutputPath
	t.Input = task.Option.InputPath
	t.Process = task.Progress.GetProgress()
	t.Status = task.Status
}
