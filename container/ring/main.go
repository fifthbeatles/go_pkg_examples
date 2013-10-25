package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// r1 0 to 9
	r1 := ring.New(10)
	for i := 0; i < 10; i++ {
		r1.Value = i
		r1 = r1.Next()
	}
	r1.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("Len:", r1.Len())

	// r2 10 to 14
	r2 := ring.New(5)
	for i := 0; i < 5; i++ {
		r2.Value = 14 - i
		r2 = r2.Prev()
	}
	r2 = r2.Move(-9)
	r2.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("Len:", r2.Len())

	// Link different Ring, r1 0 to 14
	r1 = r1.Move(9)
	r1.Link(r2)
	r1 = r1.Move(-9)
	r1.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("Len:", r1.Len())

	// Unlink, r1 5 to 14, r3 0 to 4
	r1 = r1.Move(-1)
	r3 := r1.Unlink(20)
	r1 = r1.Move(1)
	r1.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("Len:", r1.Len())
	r3.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("Len:", r3.Len())

	// Link same Ring
	r4 := r3.Move(3)
	r5 := r3.Link(r4)
	r3.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("Len:", r3.Len())
	r5.Do(func(v interface{}) {
		fmt.Printf("%v ", v)
	})
	fmt.Println("Len:", r5.Len())
}
