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
		OnDone: func(task *service.Task) {
			DefaultYouVideoCallback.CompleteChan <- task
		},
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

var codecListHandler haruka.RequestHandler = func(context *haruka.Context) {
	queryOption := service.CodecQueryOption{}
	function := context.GetQueryString("fun")
	switch function {
	case "encoder":
		queryOption.Encoding = true
	case "decoder":
		queryOption.Decoding = true
	}
	codecType := context.GetQueryString("type")
	switch codecType {
	case "video":
		queryOption.Type = service.CodecTypeVideo
	case "subtitle":
		queryOption.Type = service.CodecTypeSubtitle
	case "audio":
		queryOption.Type = service.CodecTypeAudio
	}
	search := context.GetQueryString("search")
	if len(search) > 0 {
		queryOption.Search = search
	}
	list, err := service.GetCodecList(queryOption)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	context.JSON(map[string]interface{}{
		"list": list,
	})

}

var formatListHandler haruka.RequestHandler = func(context *haruka.Context) {
	option := &service.QueryFormatOption{
		Search: context.GetQueryString("search"),
		Fun:    context.GetQueryString("fun"),
	}
	formats, err := service.ReadFormatList(option)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	context.JSON(map[string]interface{}{
		"list": formats,
	})
}
