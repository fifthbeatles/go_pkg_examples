package main

import (
	"bytes"
	"fmt"
	"os"
)

// most functions in package bytes are similar with those in package strings, not repeated

func testBuffer() {
	b := new(bytes.Buffer)
	fmt.Printf("Input Anything and ^D to finish:")
	b.ReadFrom(os.Stdin)
	if w, err := b.ReadBytes(' '); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s %d\n", w, b.Len())
	}
	b.Reset()
	b.WriteString("Hello World!")
	b.Truncate(5)
	b.WriteTo(os.Stdout)
}

func main() {
	testBuffer()
}
