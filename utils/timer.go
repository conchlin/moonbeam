package utils

import (
	"time"
)

// always use this function within a new goroutine
func CreateTimer(t time.Duration, timerAction func()) *time.Timer {
	timer := time.NewTimer(t)
	timerDone := make(chan bool)

	<-timer.C
	timerAction()
	timerDone <- true

	// block until timerDone is true
	<-timerDone

	return timer
}
