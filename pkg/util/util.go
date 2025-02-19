package util

import "time"

var NeverStop = make(chan struct{})

// TODO: add recover
func Loop(f func(), duration time.Duration, stop <-chan struct{}) {
	for range time.Tick(duration) {
		select {
		case <- stop:
			return
		default:
			f()
		}
	}
}
