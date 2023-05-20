package Id

import (
	"fmt"
)

type ConcurrentID struct {
	chanBufSize uint
	*Range
	chanId chan uint
	Close chan int
}

type Range struct {
	Start uint
	End uint
	Step uint
}

func NewCID(chanBufSize uint,_range ...Range) *ConcurrentID {
	c :=  &ConcurrentID{
		chanBufSize: chanBufSize,
		Range:getRangeSile(_range),
		chanId: make(chan uint,chanBufSize),
		Close: make(chan int),
	}

	go func() {
		go func() {
			for  {
				for i:=c.Start;i<=c.End;i+=c.Step{
					c.chanId <- i
				}
			}
		}()
		select {
		case <- c.Close:
			close(c.Close)
			close(c.chanId)
			return
		}
	}()

	return c
}

func (c *ConcurrentID) Get() uint {
	return <- c.chanId
}

func (c *ConcurrentID) Debug()  {
	go func() {
		fmt.Println(c.Get())
	}()
}