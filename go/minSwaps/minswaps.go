package main

import (
	"fmt"
	"sort"
)

func minimumSwaps(arr []int32) int32 {
	ans := 0
	var tupleList [][]int
	visited := make(map[int]bool)

	for i, el := range arr {
		pair := []int{i, int(el)}
		tupleList = append(tupleList, pair)
		visited[i] = false
	}

	sort.Slice(tupleList, func(p, q int) bool {
		return tupleList[p][1] < tupleList[q][1]
	})

	for i := 0; i < len(arr); i++ {
		// if visited or placed at correct idx continue
		if visited[i] || tupleList[i][0] == i {
			continue
		}

		cycleCount := 0
		j := i
		for !visited[j] {
			visited[j] = true

			// next node
			j = tupleList[j][0]
			cycleCount++
		}
		if cycleCount > 0 {
			ans += (cycleCount - 1)
		}
	}

	return int32(ans)
}

func main() {
	t1 := []int32{1, 5, 4, 3, 2}
	r1 := minimumSwaps(t1)
	fmt.Printf("%d \n", r1)
}
