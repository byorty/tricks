package main

import "fmt"

type Slice interface {
	Len() int
	Get(int) interface{}
}

type Ints []int

func (i Ints) Len() int {
	return len(i)
}

func (i Ints) Get(x int) interface{} {
	return i[x]
}

type Strings []string

func (s Strings) Len() int {
	return len(s)
}

func (s Strings) Get(x int) interface{} {
	return s[x]
}

func main() {
	var ints Ints = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("search", 0, ":", Search(ints, 0))
	fmt.Println("search", 5, ":", Search(ints, 5))
	fmt.Println("search", 10, ":", Search(ints, 10))

	var strings Strings = []string{"one", "two", "three"}
	fmt.Println("search", "one", ":", Search(strings, "one"))
	fmt.Println("search", "two", ":", Search(strings, "two"))
	fmt.Println("search", "some", ":", Search(strings, "four"))
}

func Search(slice Slice, x interface{}) bool {
	var hasItem bool
	for i := 0; i < slice.Len(); i++ {
		item := slice.Get(i)
		if item == x {
			hasItem = true
			break
		}
	}
	return hasItem
}
