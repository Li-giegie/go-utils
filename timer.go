package go_utils

import (
	"time"
)

func SetInterval(duration time.Duration,callFunc func())  {
	tick := time.NewTicker(duration)
	for  {
		select {
		case <- tick.C:
			callFunc()
		}
	}
}
