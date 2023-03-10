package main

import (
	"fmt"
	"github.com/gozeloglu/iterator"
)

func main() {
	iter := iterator.New(1, 2, 3, 4, 5, 6)
	for iter.Next() {
		v := iter.Value()
		fmt.Println(v)
	}

	for iter.Prev() {
		v := iter.Value()
		fmt.Println(v)
	}
	for iter.Next() {
		v := iter.Value()
		fmt.Println(v)
	}

	iterSlice := iter.ToSlice()
	fmt.Println(iterSlice)

	iter.First()     // update the cursor
	v := iter.Next() // v should be 1
	fmt.Println(v)   // prints 1

	iter.Last()
	v = iter.Next()
	fmt.Println(v) // prints 6
}
