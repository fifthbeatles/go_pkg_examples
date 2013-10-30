package main

import (
	"fmt"
	"os"
)

type st struct {
	Int       int
	Float64   float64
	String    string
	Rune      rune
	Complex64 complex64
	Byte      byte
	Bool      bool
}

func testPrint() {
	st1 := st{
		Int:       123456789,
		Float64:   12.3456789,
		String:    "亮",
		Rune:      '亮',
		Complex64: complex(12.34, 567.89),
		Byte:      'A',
		Bool:      true,
	}
	fmt.Println(st1)
	fmt.Printf("%v\n%+v\n%#v\n%%%T\n%p\n", st1, st1, st1, st1, &st1)
	fmt.Printf("%t\n", st1.Bool)
	fmt.Printf("%v %#v %T %b %d %o %x %X %U\n", st1.Int, st1.Int, st1.Int, st1.Int, st1.Int, st1.Int, st1.Int, st1.Int, st1.Int)
	fmt.Printf("%v %#v %T %b %.3e %.2E %.4f\n", st1.Float64, st1.Float64, st1.Float64, st1.Float64, st1.Float64, st1.Float64, st1.Float64)
	fmt.Printf("%v %#v %T %b %d %o %x %X %U %c %q\n", st1.Rune, st1.Rune, st1.Rune, st1.Rune, st1.Rune, st1.Rune, st1.Rune, st1.Rune, st1.Rune, st1.Rune, st1.Rune)
	fmt.Printf("%v %#v %T %s %q %x %X\n", st1.String, st1.String, st1.String, st1.String, st1.String, st1.String, st1.String)
}

func testScan() {
	var d1, d2, d3 int
	fmt.Printf("Input 3 Integers seperated by comma: ")
	if n, err := fmt.Scanf("%d,%d,%d", &d1, &d2, &d3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Got %d: %d, %d, %d\n", n, d1, d2, d3)
	}
	var s string
	var i int
	var f float64
	if n, err := fmt.Sscanf("hello 123 4.5", "%s%d%f", &s, &i, &f); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Got %d: %s, %d, %f\n", n, s, i, f)
	}
}

func testFile() {
	f, err := os.OpenFile("/tmp/fmt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Fprintln(f, "Hello World")
	fmt.Fprintln(f, 123456)
	fmt.Fprintf(f, "%d %d\n", 12, 345)
}

func main() {
	//testPrint()
	//testScan()
	testFile()
}
