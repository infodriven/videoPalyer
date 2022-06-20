package palyer

import (
	"fmt"
	"time"
)

type WMVPlayer struct {
	hourLen float32
}

//定义WMV播放器
func (p *WMVPlayer) Play(url string) {
	fmt.Println("the WMVPlayer begin play...")

	for i :=0; i <= 10; i++ {
		fmt.Printf("%d playing", i)
		time.Sleep(time.Second)
	}

	fmt.Println("the WMVPlayer end!")
}