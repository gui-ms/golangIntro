package main

import (
	"fmt"
	"time"
)

func counter(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	/*Using go routines we can easily run functions simultaneously.
	We create channels so that we can communicate between threads allocated in different memory addresses
	Sharing memory addresses is not recommended since one function can overwrite data that is being used by
	another function, rendering it useless.*/
	channel := make(chan string)

	go func() {
		channel <- "hey!"
	}()

	message := <-channel
	fmt.Println(message)
}
