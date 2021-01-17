package config

import "github.com/spf13/viper"

var DefaultConfig = ApplicationConfig{}

type ApplicationConfig struct {
	FfmpegBin  string
	FfprobeBin string
	Addr       string
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
	conf.SetDefault("addr", ":6700")
	DefaultConfig.FfmpegBin = conf.GetString("ffmpeg")
	DefaultConfig.FfprobeBin = conf.GetString("ffprobe")
	DefaultConfig.Addr = conf.GetString("addr")
	return nil
}
