package go_utils

import (
	"testing"
	"time"
)

func TestAsyncRun(t *testing.T) {
	Run(3, func() {
		time.Sleep(time.Second)
	}).Debug()
}
