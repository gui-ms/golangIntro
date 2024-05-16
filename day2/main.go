package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d \n", workerId, x)
		time.Sleep(time.Second)
	}
}

/*Using go routines we can easily run functions simultaneously.
We create channels so that we can communicate between threads allocated in different memory addresses
Sharing memory addresses is not recommended since one function can overwrite data that is being used by
another function, rendering it useless.*/

func main() {
	channel := make(chan int)

	//Here we created a load balancer with ease

	workersAmount := 10

	for j := range workersAmount {
		go worker(j, channel)
	}

	for i := range 50 {
		channel <- i
	}
}
