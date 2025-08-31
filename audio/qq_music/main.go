package main

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 环境：
// ffmpeg
// go get github.com/u2takey/ffmpeg-go

// 将QQ音乐的ogg文件转化为mp3文件
func main() {
	// ffmpeg -y -i input.ogg -acodec libmp3lame -ab 320k output.mp3

	//设置音频比特率（audio bitrate）
	//128k → 普通音质（够用）
	//192k → 中高音质
	//320k → 高音质（接近无损，但文件大）
	ab := "320k"
	err := ffmpeg.Input("music.ogg").
		Output("music.mp3", ffmpeg.KwArgs{"ab": ab, "acodec": "libmp3lame"}).
		OverWriteOutput().
		Run()
	if err != nil {
		fmt.Println("音乐转化出错！", err)
	}
}
