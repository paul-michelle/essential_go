package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne, channelTwo := make(chan int), make(chan int)

	go func() {
		channelOne <- 345
	}()

	select {
	case value := <-channelOne:
		fmt.Printf("received from channel 1: %d\n", value)
	case value := <-channelTwo:
		fmt.Printf("received from channle 2: %d\n", value)

	}

	channelThree := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		value := 3.14
		channelThree <- value
	}()

	select {
	case value:= <- channelThree:
		fmt.Printf("received from channel 3: %f\n", value)
	case <- time.After(101 * time.Millisecond):
		fmt.Println("timeout!")
	}
	
}
