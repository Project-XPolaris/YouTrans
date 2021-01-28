package service

import (
	"github.com/ahmetb/go-linq/v3"
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

type CodecType int

const (
	CodecTypeVideo = iota + 1
	CodecTypeAudio
	CodecTypeSubtitle
)

type CodecQueryOption struct {
	Decoding bool
	Encoding bool
	Type     CodecType
}

func GetCodecList(queryOption CodecQueryOption) ([]ffmpeg.Codec, error) {
	conf := GetConfig()
	list, err := ffmpeg.ReadCodecList(conf)
	if err != nil {
		return nil, err
	}
	query := linq.From(list)
	if queryOption.Encoding {
		query = query.Where(func(i interface{}) bool {
			return i.(ffmpeg.Codec).Flags.Encoding
		})
	}
	if queryOption.Decoding {
		query = query.Where(func(i interface{}) bool {
			return i.(ffmpeg.Codec).Flags.Decoding
		})
	}
	switch queryOption.Type {
	case CodecTypeVideo:
		query = query.Where(func(i interface{}) bool {
			return i.(ffmpeg.Codec).Flags.VideoCodec
		})
	case CodecTypeAudio:
		query = query.Where(func(i interface{}) bool {
			return i.(ffmpeg.Codec).Flags.AudioCodec
		})
	case CodecTypeSubtitle:
		query = query.Where(func(i interface{}) bool {
			return i.(ffmpeg.Codec).Flags.SubtitleCodec
		})
	}
	query.ToSlice(list)
	return list, nil
}
