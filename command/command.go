package command

import (
	"fmt"
	"strconv"

	manager "github.com/infodriven/videoPalyer/manager"
	play "github.com/infodriven/videoPalyer/play"
)

//全局资源管理器
var videoManager *manager.VideoManager = manager.NewVideoManager()

func HandleComand(parms []string) {
	switch parms[0] {
	case "add":
		videoType, _ := strconv.Atoi(parms[3])
		videoManager.AddVideo(parms[1], parms[2], videoType)
	case "del":
		videoManager.DelVideo(parms[1])
	case "list":
		videolist := videoManager.ListVideo()
		for i, v := range videolist {
			fmt.Printf("%d: Id=%d Name=%s Url=%s VideoType=%d\n", i, v.ID, v.Name, v.Url, v.VideoType)
		}
	case "play":
		videoEntry := videoManager.FindName(parms[1])
		play.PalyVideo(videoEntry.Url, videoEntry.VideoType)
	default:
		fmt.Println("not found command")

	}
}
