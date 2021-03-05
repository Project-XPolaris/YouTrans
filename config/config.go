package config

import "github.com/spf13/viper"

var DefaultConfig = ApplicationConfig{}

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
	DefaultConfig.FfmpegBin = conf.GetString("ffmpeg")
	DefaultConfig.FfprobeBin = conf.GetString("ffprobe")
	DefaultConfig.Addr = conf.GetString("addr")
	DefaultConfig.YouVideoUrl = conf.GetString("youvideourl")
	DefaultConfig.YouvideoEnable = conf.GetBool("youvideoenable")
	return nil
}
