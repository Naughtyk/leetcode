/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked (the number I picked stays the same throughout the game).

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(guessNumber(10))
}

func guess(num int) int {
	pick := 6
	if num < pick {
		return 1
	}
	if num == pick {
		return 0
	}
	return -1
}

func guessNumber(n int) int {
	left := 1
	right := n
	for {
		middle := (right + left) / 2
		if guess(middle) == 0 {
			return middle
		} else if guess(middle) == -1 {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
}
