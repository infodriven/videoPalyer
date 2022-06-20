package palyer

import (
	"fmt"
	"time"
)

type MP4Player struct {
	hourLen float32
}

//定义MP4播放器
func (p *MP4Player) Play(url string) {
	fmt.Println("the MP4Player begin play...")

	for i :=0; i <= 10; i++ {
		fmt.Printf("%d playing\n", i)
		time.Sleep(time.Second)
	}

	fmt.Println("the MP4Player end!")
}
