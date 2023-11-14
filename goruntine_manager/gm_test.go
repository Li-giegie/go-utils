package goruntine_manager

import (
	"fmt"
	"log"
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
		for i := 0; i < 10; i++ {
			n := i
			if err := gm.Dilatation(1); err != nil {
				log.Fatalln(err)
			}
			gm.Exit()
			err := gm.Run(func() {
				time.Sleep(time.Second)
				fmt.Println(n)
			})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()

	select {}
}
