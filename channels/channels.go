package channels

import "fmt"

const Attempts = 7

func receiver(ch chan int) {
	for range [Attempts]struct{}{} {
		// set the condition in this function to stop the loop in time
		if value, ok := <-ch; ok {
			fmt.Println(value)
		}
	}
}

func sender(send int) chan int {
	ch := make(chan int, send)

	go func() {
		defer close(ch)
		for i := 0; i < send; i++ {
			ch <- i
		}
	}()
	return ch
}

func RunConcurrentAttempts() {
	var send int
	fmt.Scan(&send)

	ch := sender(send)
	receiver(ch)
}
