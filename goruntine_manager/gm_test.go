package goruntine_manager

import (
	"fmt"
	"testing"
	"time"
)

func TestNewGoroutineManger(t *testing.T) {
	gm := NewGoroutineManger(1)
	if err := gm.Start(); err != nil {
		t.Error(err)
		return
	}
	go func() {
		gm.Run(func() {
			for {
				fmt.Println("running")
				time.Sleep(time.Second)
			}
		})
	}()
	time.Sleep(time.Second)
	gm.Exit()
	fmt.Println("exit")
	select {}
}
