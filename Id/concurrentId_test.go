package Id

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCId(t *testing.T) {
	id := NewCID(100,Range{
		Start: 1,
		End:   5,
		Step:  2,
	})
	n:= 100000
	t1 := time.Now()
	for i:=0;i<n;i++{
		id.Get()
	}

	fmt.Println(n,"次数	耗时",time.Since(t1))
}
