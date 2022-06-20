package mananger

import "time"

//定义视频管理器
type VideoManager struct {
	videoList []VideoEntry
}

//对外暴露的使用接口，用于获取视屏管理器的句柄
func NewVideoManager() *VideoManager {
	vm := &VideoManager{
		videoList: make([]VideoEntry, 0, 128),
	}

	return vm
}

//添加视频
func (vm *VideoManager) AddVideo(name string, url string, videoType int) {
	ve := VideoEntry{
		ID:        int(time.Now().UnixNano()),
		Name:      name,
		Url:       url,
		VideoType: videoType,
	}

	vm.videoList = append(vm.videoList, ve)
}

//删除视频
func (vm *VideoManager) DelVideo(name string) *VideoEntry {
	index := -1 //要删除的视频在切片中的索引
	for i, v := range vm.videoList {
		if v.Name == name {
			index = i
			break
		}
	}
	if index == -1 {
		return nil
	}

	rst := &vm.videoList[index]

	vm.videoList = append(vm.videoList[0:index], vm.videoList[index+1:]...)    //三个点是打散的意思

	return rst
}

//显示所有的视频列表
func (vm *VideoManager) ListVideo() []VideoEntry {
	return vm.videoList
}

//根据name获取视频实体
func (vm *VideoManager) FindName(name string) *VideoEntry{
	index := -1
	for i, v := range vm.videoList {
		if v.Name == name {
			index = i
		}
	}
	
	if index == -1 {
		return nil
	}

	return &vm.videoList[index]
}