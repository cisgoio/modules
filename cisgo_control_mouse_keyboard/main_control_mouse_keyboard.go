package main

import (
	"cisgo_control_mouse_keyboard/robotgo"
	"fmt"
	"time"
)

func main() {

	for {
		dt := time.Now()
		fmt.Println("sk: keep mouse moving, sleep 75500, v.13.4 , Current date/time : ", dt.String())

		robotgo.MouseSleep = 75500

		robotgo.ScrollMouse(1, "up")
		//robotgo.ScrollMouse(20, "right")
		//robotgo.DragSmooth(1, 1)

	}

	/*
		    //==OTHER Possibilities :-D
			robotgo.Scroll(0, -10)
			robotgo.Scroll(100, 0)
			robotgo.MilliSleep(100)
			robotgo.ScrollSmooth(-10, 6)
			// robotgo.ScrollRelative(10, -100)
			robotgo.Move(10, 20)
			robotgo.MoveRelative(0, -10)
			robotgo.DragSmooth(10, 10)
			robotgo.Click("wheelRight")
			robotgo.Click("left", true)
			robotgo.MoveSmooth(100, 200, 1.0, 10.0)
			robotgo.Toggle("left")
			robotgo.Toggle("left", "up")
	*/
}
