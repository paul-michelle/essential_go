// classic modulo devision

package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {

	fizzInt := askForInt("fizz")
	buzzInt := askForInt("buz")

	for i:=1; i<21; i++ {

		if i % fizzInt * buzzInt == 0 {
			fmt.Println("fizzbuzz")
			continue
		}

		if i % fizzInt == 0 {
			fmt.Println("fizz")
			continue
		}

		if i % buzzInt == 0 {
			fmt.Println("buzz")
			continue
		}

		fmt.Println(i)
	}
}

func askForInt(exclamation string) int {
	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)

	fmt.Printf("Enter an integer for '%s': ", exclamation)
	valueEntered, _ := reader.ReadString('\n')
	valueEnteredTrimmed := strings.TrimSpace(valueEntered)
	intEntered, error := strconv.ParseInt(valueEnteredTrimmed, 10 ,32)
	if error != nil {
		panic(error)
	}
	return int(intEntered)
}