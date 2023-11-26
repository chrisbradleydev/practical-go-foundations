package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) { // use parameter
			defer wg.Done()
			fmt.Println(n)
		}(i)
		/* "i" shadows "i" from the for loop
		i := i
		go func() {
			fmt.Println(i)
		}()
		*/
		/* bug: all goroutines use final "i" in the for loop
		go func() {
			fmt.Println(i)
		}()
		*/
	}
	wg.Wait()

	// time.Sleep(10 * time.Millisecond) // don't use sleep

	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	msg := <- ch // receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("received", msg)
	}

	msg = <- ch
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <- ch
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)

	// ch <- "hi" // channel is closed -> panic
	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))
}

// the worst sort algorithm possible
func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		n := n
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}

	var out []int
	// for i = 0; i < len(values); i++ {}
	for range values {
		n := <- ch
		out = append(out, n)
	}
	return out
}

/*
channel semantics
- send & receive will block until opposite operation (*)
- receive from a closed channel will return the zero value without blocking
- send to a closed channel will panic
- closing a closed channel will panic
- send/receive to a nil channel will block forever
*/