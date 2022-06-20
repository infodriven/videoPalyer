package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/infodriven/videoPalyer/command"
)


func main() {
	fmt.Println("VideoPlayer start ...welcome to you!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please input you choice:")

		raw, _, _ := reader.ReadLine()
		line := string(raw)

		if line == "quit" {
			fmt.Println("goodby!")
			break
		}
		parms := strings.Split(line, " ")

		command.HandleComand(parms)
	}	
}

