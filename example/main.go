package main

import (
	"fmt"
	"github.com/gozeloglu/iterator"
)

func main() {
	iter := iterator.New(1, 2, 3, 4, 5, 6)
	for iter.HasNext() {
		v := iter.Next()
		fmt.Println(v)
	}

	for iter.HasPrev() {
		v := iter.Prev()
		fmt.Println(v)
	}
}
