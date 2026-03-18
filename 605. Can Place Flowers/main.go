/*
You have a long flowerbed in which some of the plots are planted, and some are not. However,
 flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
 return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	if len(flowerbed) == 1 {
		if n == 0 {
			return true
		}
		return flowerbed[0] == 0
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		n--
		flowerbed[0] = 1
	}

	i := 1
	for i < len(flowerbed)-1 {
		if n == 0 {
			return true
		}
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			n--
			i++
		}
		i++
	}
	if i < len(flowerbed) && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
		n--
	}
	if n <= 0 {
		return true
	}
	return false
}
