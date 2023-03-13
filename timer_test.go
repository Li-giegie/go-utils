package go_utils

import (

	"testing"
	"time"
)

//
func Test_SetInterval(t *testing.T) {
	SetInterval(time.Second*2, func() {
		println("SetInterval------")
	})
}
//
func Test_SetTimer(t *testing.T) {
	SetTimer(time.Second, func() {
		println("Test_SetTimer------")
	})
}