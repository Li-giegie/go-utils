package go_utils

import (

	"testing"
	"time"
)


func Test_SetInterval(t *testing.T) {
	SetInterval(time.Second*1, func() {
		println("asdasd")
	})
}
