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

type BankAccount struct {
	RWmutex sync.RWMutex
	Balance int
}

// method to add balance
func (account *BankAccount) AddBalance(amount int) {
	account.RWmutex.Lock() //lock for writing
	account.Balance += amount
	account.RWmutex.Unlock() //unlock for writing
}

// method to get balance
func (account *BankAccount) GetBalance() int {
	account.RWmutex.RLock() //lock for reading
	balance := account.Balance
	account.RWmutex.RUnlock() //unlock for reading

	return balance
}

func TestRwMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {

				account.AddBalance(1)
				fmt.Println("current balance", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("final balance", account.GetBalance())
}
