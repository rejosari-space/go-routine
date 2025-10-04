package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(id int, wg *sync.WaitGroup) {
	defer wg.Done() //must be done, to tell that this goroutine is done

	wg.Add(1) //add the counter of wait group

	fmt.Println("hello from goroutine", id)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 1; i < 100; i++ {
		go RunAsync(i, wg)

	}

	wg.Wait() //wait until the counter of wait group is 0

	fmt.Println("all goroutine is done")
}
