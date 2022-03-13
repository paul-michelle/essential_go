package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			channel <- i
			fmt.Printf("Channel got %d\n", i)
			time.Sleep(time.Second)
		}
		close(channel)
	}()

	for value := range channel {
		fmt.Printf("Received from channel %d\n", value)
	}
	
}
