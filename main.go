package main

import (
	"fmt"
	"time"
)

func main() {

	//example1()
	//example2()
	//example3()

	// worker queue example1
	// bufgfer channels of size 100
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	//
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// example 3 with func count
func example3() {

	// channel send and receive is blocking function
	c := make(chan string)
	go count("sheep", c)
	for {
		//
		msg, open := <-c // receive value from the channel, if channel is empty, blocking happen
		// if channel is closed, break the loop
		if !open {
			break
		}
		fmt.Println(msg)
	}
}

// go functions
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing // send value to channel, see it as return channel of c
		time.Sleep(time.Millisecond * 500)
	}
	close(c) // as a sender we can close the channel
}

// example1 both goroutine will execute and wait on each other, even if the first
// one is four times faster. We have to use select (example2)
func example1() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)

	}
}

// example2 will use select to process that channel that are ready to receive
func example2() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	// will process the channel that are ready to receive
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// worker documentation, only receive on jobs channel,
// only send on results channel. Otherwise compile error
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

// fib documentation
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
