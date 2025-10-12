package Timer

import (
	"fmt"
	"time"
)

func StartTimer() func() {
	start := time.Now()
	return func() {
		elapsed := float64(time.Since(start).Microseconds()) / 1000
		fmt.Printf("Elapsed: %v ms\n", elapsed)
	}
}
