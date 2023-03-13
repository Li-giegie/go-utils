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

func SetTimer(duration time.Duration,callFunc func()){
	timer := time.NewTimer(duration)
	select {
	case <- timer.C:
		callFunc()
	}
}
