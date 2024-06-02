package main

import (
	"flag"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	toggle := flag.Bool("t", false, "toggle, used for keeping the jiggler on, if set to true putting the mouse in the top-left coorner will toggle the jiggling")
	isJiggling := true
	flag.Parse()
	robotgo.MouseSleep = 100
	sx, _ := robotgo.GetScreenSize()
	for {
		if isJiggling {
			robotgo.MoveRelative(0, 1)
			robotgo.MoveRelative(0, -1)
		}
		time.Sleep(time.Second)
		locX, locY := robotgo.Location()
		if locX >= sx-8 && locY <= 8 {
			if !*toggle {
				break
			} else {
				isJiggling = !isJiggling
				time.Sleep(time.Second * 3)
			}
		}
	}
}
