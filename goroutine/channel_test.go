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

// set capacity as temporary storage for channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "abdul"
	channel <- "aziz"
	channel <- "ganteng"

	fmt.Println("data from channel =>", <-channel)
	fmt.Println("data from channel =>", <-channel)

}

// with goroutine
func TestBufferedChannelGoroutine(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// sender
	go func() {
		channel <- "abdul"
		channel <- "aziz"
		channel <- "ganteng banget"

	}()

	// receiver
	go func() {

		fmt.Println("data from channel =>", <-channel)
		fmt.Println("data from channel =>", <-channel)
		fmt.Println("data from channel =>", <-channel)

	}()

	time.Sleep(2 * time.Second)

	fmt.Println("process was done 👋")

}
