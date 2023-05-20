package id

import (
	"fmt"
	"math"
	"os"
)

func getRangeSile(r []Range) *Range {
	if len(r) == 0 {
		return &Range{
			Start: 0,
			End:   math.MaxUint,
			Step:  1,
		}
	}
	if r[0].Start > r[0].End {
		fmt.Println("Range err : start is greater than end")
		os.Exit(1)
	}
	if r[0].Step > r[0].End {
		fmt.Println("Range err : step is greater than end")
		os.Exit(1)
	}
	return &r[0]
}
