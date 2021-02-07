package service

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/allentom/transcoder/ffmpeg"
	"github.com/projectxpolaris/youtrans/config"
	"strings"
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
	Search   string
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
	if len(queryOption.Search) > 0 {
		query = query.Where(func(i interface{}) bool {
			return strings.Contains(i.(ffmpeg.Codec).Name, queryOption.Search) || strings.Contains(i.(ffmpeg.Codec).Desc, queryOption.Search)
		})
	}
	query.ToSlice(&list)
	return list, nil
}

type QueryFormatOption struct {
	Search string
	Fun    string
}

func ReadFormatList(option *QueryFormatOption) ([]ffmpeg.SupportFormat, error) {
	conf := GetConfig()
	list, err := ffmpeg.GetFormats(conf)
	if err != nil {
		return nil, err
	}
	query := linq.From(list)
	if len(option.Search) > 0 {
		query = query.Where(func(i interface{}) bool {
			return strings.Contains(i.(ffmpeg.SupportFormat).Name, option.Search) || strings.Contains(i.(ffmpeg.SupportFormat).Desc, option.Search)
		})
	}

	if len(option.Fun) > 0 {
		switch option.Fun {
		case "mux":
			query = query.Where(func(i interface{}) bool {
				return i.(ffmpeg.SupportFormat).Flags.Muxing
			})
		case "dmux":
			query = query.Where(func(i interface{}) bool {
				return i.(ffmpeg.SupportFormat).Flags.Demuxing
			})
		}

	}

	query.ToSlice(&list)
	return list, nil
}
