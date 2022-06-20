package palyer

import "fmt"

//定义播放接口
type IPlay interface {
	Play(Url string)
}

func PalyVideo(url string, videoType int) {
	var p IPlay

	switch videoType {
	case 1:
		p = &MP4Player{2.5}
	case 2:
		p = &WMVPlayer{1.2}
	default:
		fmt.Println("not found the new videotype")
		return
	}

	p.Play(url)
}
