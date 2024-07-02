/*
You have n jobs and m workers. You are given three arrays:
 difficulty, profit, and worker where:

difficulty[i] and profit[i] are the difficulty and the profit of the ith job, and
worker[j] is the ability of jth worker (i.e., the jth worker can only
	 complete a job with difficulty at most worker[j]).
Every worker can be assigned at most one job, but one job can
 be completed multiple times.

For example, if three workers attempt the same job that pays $1,
 then the total profit will be $3. If a worker cannot complete
  any job, their profit is $0.
Return the maximum profit we can achieve after assigning the
 workers to the jobs.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProfitAssignment([]int{2, 4, 6, 8, 10},
		[]int{10, 20, 30, 40, 50},
		[]int{4, 5, 6, 7}))
}

type work struct {
	diff int
	prof int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var listWork []work
	for i := range len(difficulty) {
		listWork = append(listWork, work{diff: difficulty[i], prof: profit[i]})
	}
	sort.Slice(listWork, func(i, j int) bool {
		return listWork[i].diff < listWork[j].diff
	})
	sort.Ints(worker)
	ret := 0
	currMax := 0
	j := 0
	for _, empl := range worker {
		for j < len(listWork) && empl >= listWork[j].diff {
			currMax = max(currMax, listWork[j].prof)
			j++
		}
		ret += currMax
	}

	return ret
}
