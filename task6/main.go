package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("nil channel")
	testNilChannel()

	fmt.Println("closed channel")
	testClosedChannel()

	fmt.Println("not-nil & not-closed channel")
	testNormalChannel()
}

func testNilChannel() {
	var nilChan chan int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("   [NilChannel] close() вызвал panic:", r)
			}
		}()
		close(nilChan)
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("   [NilChannel] send() вызвал panic:", r)
			}
		}()
		nilChan <- 42
	}()

	go func() {
		done := make(chan struct{})
		go func() {
			defer close(done)
			_ = <-nilChan
		}()
		select {
		case <-done:
		case <-time.After(1 * time.Second):
			fmt.Println("   [NilChannel] receive() блокируется навсегда")
		}
	}()

	time.Sleep(100 * time.Millisecond)
}

func testClosedChannel() {
	closedChan := make(chan int, 1)
	closedChan <- 10
	close(closedChan)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("   [ClosedChannel] close() вызвал panic:", r)
			}
		}()
		close(closedChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("   [ClosedChannel] send() вызвал panic:", r)
			}
		}()
		closedChan <- 42
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		v, ok := <-closedChan
		fmt.Printf("   [ClosedChannel] получено: %v, ok=%v\n", v, ok)
	}()

	wg.Wait()
}

func testNormalChannel() {
	normalChan := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		normalChan <- 100
		fmt.Println("   [NormalChannel] отправлено: 100")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case v := <-normalChan:
			fmt.Printf("   [NormalChannel] получено: %v\n", v)
		case <-time.After(1 * time.Second):
			fmt.Println("   [NormalChannel] receive() блокируется")
		}
	}()

	wg.Wait()
	close(normalChan)
	fmt.Println("   [NormalChannel] канал успешно закрыт")

	normalChan2 := make(chan int, 1)
	normalChan2 <- 200
	close(normalChan2)
	v, ok := <-normalChan2
	fmt.Printf("   [NormalChannel] после закрытия: %v, ok=%v\n", v, ok)
}
