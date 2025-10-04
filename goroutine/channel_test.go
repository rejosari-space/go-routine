package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func GivmeResponse(channel chan string) {

	time.Sleep(2 * time.Second)
	channel <- "this is response data"
}

// func channel can only send data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "this is response data"
}

// func cannel only receiove data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("data from channel =>", data)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel) //must be close

	go GivmeResponse(channel)

	data := <-channel

	fmt.Println("data from channel =>", data)

	time.Sleep(5 * time.Second)
}

func TestChannelOnlyInAndOut(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)

}
