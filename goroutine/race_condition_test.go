package goroutine

import (
	"fmt"
	"testing"
	"time"
)

// its dangerous if multiple goroutine access the same variable
// because the result will be unpredictable

// so we can handle with mutex
func TestRaceCondition(t *testing.T) {

	x := 0

	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()

	}

	time.Sleep(5 * time.Second)
	fmt.Println("value of =>", x)
}
