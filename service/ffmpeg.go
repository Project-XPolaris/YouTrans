package service

import (
	"github.com/allentom/transcoder/ffmpeg"
	"github.com/projectxpolaris/youtrans/config"
)

func GetConfig() *ffmpeg.Config {
	FfmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   config.DefaultConfig.FfmpegBin,
		FfprobeBinPath:  config.DefaultConfig.FfprobeBin,
		ProgressEnabled: true,
	}
	return FfmpegConf
}
