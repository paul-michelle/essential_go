package main

import (
	"fmt"
)

func getValueAt(values []int, index int) int {
	fmt.Println("getValueAt func started.")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error occurred: %s\n", err)
		}
	}()
	fmt.Println("getValueAt func about to return.")
	return values[index]
}

func main() {
	fmt.Println("main func entered.")

	values := []int{1, 2, 3}

	fmt.Println("About to call getValueAt func from main func.")

	value := getValueAt(values, 4)

	fmt.Println("Got result of getValueAt execution. About to print it out.")

	fmt.Println(value)
}
