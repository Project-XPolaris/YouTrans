package config

import "github.com/spf13/viper"

var Instance = ApplicationConfig{}

type ApplicationConfig struct {
	FfmpegBin      string
	FfprobeBin     string
	Addr           string
	YouVideoUrl    string
	YouvideoEnable bool
}

func LoadConfig() error {
	conf := viper.New()
	conf.AddConfigPath("./")
	conf.AddConfigPath("../")
	conf.SetConfigName("config")
	conf.SetConfigType("yaml")

	err := conf.ReadInConfig()
	if err != nil {
		return err
	}
	conf.SetDefault("ffmpeg", "/usr/bin/ffmpeg")
	conf.SetDefault("ffprobe", "/usr/bin/ffprobe")
	conf.SetDefault("youvideo_enable", false)
	conf.SetDefault("addr", ":6700")
	Instance.FfmpegBin = conf.GetString("ffmpeg")
	Instance.FfprobeBin = conf.GetString("ffprobe")
	Instance.Addr = conf.GetString("addr")
	Instance.YouVideoUrl = conf.GetString("youvideourl")
	Instance.YouvideoEnable = conf.GetBool("youvideoenable")
	return nil
}
