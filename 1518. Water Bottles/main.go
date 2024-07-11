/*
There are numBottles water bottles that are initially full of water. You can exchange numExchange empty water bottles from the market with one full water bottle.

The operation of drinking a full water bottle turns it into an empty bottle.

Given the two integers numBottles and numExchange, return the maximum number of water bottles you can drink.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(numWaterBottles(9, 3) == 13)
}

func numWaterBottles(numBottles int, numExchange int) int {
	ret := 0
	empty := 0
	for {
		ret += numBottles
		numBottles += empty
		empty = numBottles % numExchange
		numBottles = numBottles / numExchange
		if numBottles == 0 {
			return ret
		}
	}
}
