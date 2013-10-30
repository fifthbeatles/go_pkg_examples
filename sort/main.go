package main

import (
	"fmt"
	"sort"
)

func sortInt() {
	s := []int{3, 9, 4, 8, 7, 2, 6, 1, 0}
	fmt.Println(s, sort.IntsAreSorted(s))
	sort.Ints(s)
	fmt.Println(s, sort.IntsAreSorted(s))
	// SearchXXXs return the position to insert the new value, s must be sorted
	fmt.Println(sort.SearchInts(s, -1), sort.SearchInts(s, 4), sort.SearchInts(s, 5), sort.SearchInts(s, 10))
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}

// sort Sort.Interface
type Organ struct {
	Name   string
	Weight int
}

type Organs []*Organ

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByName struct {
	Organs
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

type ByWeight struct {
	Organs
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func sortInterface() {
	s := []*Organ{
		&Organ{"brain", 1340},
		&Organ{"heart", 290},
		&Organ{"liver", 1494},
		&Organ{"pancreas", 131},
		&Organ{"prostate", 62},
		&Organ{"spleen", 162},
	}

	sort.Sort(ByName{s})
	fmt.Println("Order by Name")
	printOrgans(s)

	sort.Sort(ByWeight{s})
	fmt.Println("Order by Weight")
	printOrgans(s)

	f := func(search int) (f func(i int) bool) {
		return func(i int) bool {
			return s[i].Weight > search
		}
	}

	fmt.Println(sort.Search(len(s), f(300)), sort.Search(2, f(300)))
}

func printOrgans(s []*Organ) {
	for _, o := range s {
		fmt.Printf("%s:%d\n", o.Name, o.Weight)
	}
}

func main() {
	sortInt()
	sortInterface()
}
