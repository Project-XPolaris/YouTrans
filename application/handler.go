package application

import (
	"github.com/allentom/haruka"
	"github.com/allentom/transcoder/ffmpeg"
	"github.com/projectxpolaris/youtrans/service"
	"net/http"
)

type CreateTaskRequestBody struct {
	Input  string `json:"input"`
	Output string `json:"output"`
	Format string `json:"format"`
	Codec  string `json:"codec"`
}

var createTransHandler haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody CreateTaskRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	opts := service.TaskOption{
		Option: ffmpeg.Options{
			VideoCodec: &requestBody.Codec,
		},
		Overwrite:  true,
		Format:     requestBody.Format,
		InputPath:  requestBody.Input,
		OutputPath: requestBody.Output,
	}
	err = service.DefaultTaskPool.CreatTask(&opts)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	context.JSON(map[string]interface{}{
		"success": true,
	})
}

var taskListHandler haruka.RequestHandler = func(context *haruka.Context) {
	tasks := make([]TaskTemplate, 0)
	for _, task := range service.DefaultTaskPool.Tasks {
		template := TaskTemplate{}
		template.Assign(task)
		tasks = append(tasks, template)
	}
	context.JSON(map[string]interface{}{
		"success": true,
		"list":    tasks,
	})
}

var stopTransHandler haruka.RequestHandler = func(context *haruka.Context) {
	id := context.GetQueryString("id")
	service.DefaultTaskPool.StopTask(id)
	context.JSON(map[string]interface{}{
		"success": true,
	})
}
