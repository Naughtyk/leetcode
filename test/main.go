package main

import (
	"fmt"
)

func main() {

	var numbers []*int
	for _, v := range []int{10, 20, 30, 40} {
		numbers = append(numbers, &v)
	}

	for _, num := range numbers {
		fmt.Printf("%d ", *num)
	}

}
