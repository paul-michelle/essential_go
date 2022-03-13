// interfaces, io.Writer
package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	wtr io.Writer
}

func NewCapper(wtr io.Writer) *Capper {
	capper := &Capper{wtr: wtr}
	return capper
}

func (c *Capper) Write(text []byte) (n int, err error) {
	unicodeDistance := byte('a' - 'A')

	output := make([]byte, len(text))

	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			output = append(output, char - unicodeDistance)
			continue
		}
		output = append(output, char)

	}
	return c.wtr.Write(output)
}

func main() {
	capper := NewCapper(os.Stdout)
	fmt.Fprintln(capper, "capper has been used")
}