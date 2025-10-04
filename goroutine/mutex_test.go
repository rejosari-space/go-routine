package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var mutex sync.Mutex
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {

			for j := 1; j <= 100; j++ {
				mutex.Lock() //lock the variable x
				x += 1
				mutex.Unlock() //unlock the variable x
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("value of x =>", x)
}
