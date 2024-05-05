package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.MouseSleep = 100
	sx, _ := robotgo.GetScreenSize()
	for {
		robotgo.Move(sx, 10)
		robotgo.Move(sx, 11)
		time.Sleep(time.Second)
		locX, locY := robotgo.Location()
		if locX >= sx-8 && locY <= 8 {
			break
		}
	}
}
