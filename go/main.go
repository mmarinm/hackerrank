package main

import "fmt"

func main() {

	// minswaps
	t1 := []int32{1, 5, 4, 3, 2}
	r1 := minimumSwaps(t1)
	fmt.Printf("%d \n", r1)

	//new year chaos
	// input := []int32{2, 1, 5, 3, 4}
	// input2 := []int32{2, 5, 1, 3, 4}
	input := []int32{5, 1, 2, 3, 7, 8, 6, 4}
	input2 := []int32{1, 2, 5, 3, 4, 7, 8, 6}
	minimumBribes(input)
	minimumBribes(input2)

}
