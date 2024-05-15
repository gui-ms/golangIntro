package main

import (
	"fmt"
	"time"
)

func contador(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	//Using go routines we can easily run these functions simultaneously
	go contador(10)
	go contador(10)
	contador(10)
}
