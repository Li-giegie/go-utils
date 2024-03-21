package go_utils

import (
	"fmt"
	"sync"
	"time"
)

type TResult struct {
	N    int
	Avg  time.Duration
	Sum  time.Duration
	Mode string
}

func newTResult(start, end time.Time, n int, m string) *TResult {
	tr := new(TResult)
	tr.N = n
	tr.Sum = end.Sub(start)
	tr.Avg = tr.Sum / time.Duration(n)
	tr.Mode = m
	return tr
}

func (t *TResult) String() string {
	return fmt.Sprintf("result: sum duration: [%v], avg time: [%v], num: [%v], mode: [%v],", t.Sum.String(), t.Avg.String(), t.N, t.Mode)
}

func (t *TResult) Debug() {
	fmt.Println(t.String())
}

func AsyncRun(num int, f func()) *TResult {
	var w sync.WaitGroup
	w.Add(num)
	start := time.Now()
	for i := 0; i < num; i++ {
		go func() {
			defer w.Done()
			f()
		}()
	}
	w.Wait()
	end := time.Now()
	return newTResult(start, end, num, "AsyncRun")
}

func Run(num int, f func()) *TResult {
	start := time.Now()
	for i := 0; i < num; i++ {
		f()
	}
	end := time.Now()
	return newTResult(start, end, num, "Sync")
}

func AsyncRunWithIndex(num int, f func(i int)) *TResult {
	var w sync.WaitGroup
	w.Add(num)
	start := time.Now()
	for i := 0; i < num; i++ {
		go func(i int) {
			defer w.Done()
			f(i)
		}(i)
	}
	w.Wait()
	end := time.Now()
	return newTResult(start, end, num, "AsyncRun")
}
