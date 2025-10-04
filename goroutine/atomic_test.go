package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// atomic package is used to handle variable in goroutine without mutex

func TestAtomic(t *testing.T) {
	var x int64 = 0
	wg := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				// x += 1
				// use atomic package to handle variable in goroutine
				atomic.AddInt64(&x, 1)
			}
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("value of x =>", x)

}
