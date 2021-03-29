package application

import (
	"github.com/allentom/transcoder/ffmpeg"
	"github.com/projectxpolaris/youtrans/service"
)

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

type BaseCodecTemplate struct {
	Name string   `json:"name"`
	Desc string   `json:"desc"`
	Type string   `json:"type"`
	Feat []string `json:"feat"`
}

func (t *BaseCodecTemplate) Serializer(dataModel interface{}, context map[string]interface{}) error {
	model := dataModel.(ffmpeg.Codec)
	t.Name = model.Name
	t.Desc = model.Desc
	if model.Flags.AudioCodec {
		t.Type = "Audio"
	}
	if model.Flags.VideoCodec {
		t.Type = "Video"
	}
	if model.Flags.SubtitleCodec {
		t.Type = "Subtitle"
	}
	t.Feat = []string{}
	if model.Flags.Decoding {
		t.Feat = append(t.Feat, "decode")
	}
	if model.Flags.Encoding {
		t.Feat = append(t.Feat, "encode")
	}
	return nil
}

type BaseFormatTemplate struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (t *BaseFormatTemplate) Serializer(dataModel interface{}, context map[string]interface{}) error {
	model := dataModel.(ffmpeg.SupportFormat)
	t.Name = model.Name
	t.Desc = model.Desc
	return nil
}
