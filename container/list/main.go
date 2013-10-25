package main

import (
	"container/list"
	"fmt"
)

func main() {
	l1 := list.New()
	e1 := l1.PushFront(1)
	e4 := l1.PushBack(4)
	l1.InsertBefore(3, e4)
	l1.InsertAfter(2, e1)
	ex := l1.InsertBefore(0.2, e4)
	ey := l1.InsertAfter("Hello", e1)
	for e := l1.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(int); ok {
			fmt.Printf("%d ", v)
		} else {
			fmt.Printf("(%v) ", e.Value)
		}
	}
	fmt.Println()
	fmt.Println("Len:", l1.Len())
	for e := l1.Back(); e != nil; e = e.Prev() {
		if v, ok := e.Value.(int); ok {
			fmt.Printf("%d ", v)
		} else {
			fmt.Printf("(%v) ", e.Value)
		}
	}
	fmt.Println()
	fmt.Println("Len:", l1.Len())
	l1.MoveToBack(e1)
	l1.MoveToFront(e4)
	l1.Remove(ex)
	l1.Remove(ey)
	for e := l1.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
	fmt.Println("Len:", l1.Len())
	l2 := list.New()
	l2.PushBack(10)
	l1.PushBackList(l2)
	l2.PushFront(5)
	l1.PushFrontList(l2)
	for e := l1.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
	fmt.Println("Len:", l1.Len())
	l1.Init()
	fmt.Println("Len:", l1.Len())
}
