package main

import "fmt"

func main() {
	fmt.Printf("minswaps \n")
	// minswaps
	t1 := []int32{1, 5, 4, 3, 2}
	r1 := minimumSwaps(t1)
	fmt.Printf("%d \n", r1)

	fmt.Printf("new year chaos \n")
	//new year chaos
	// input := []int32{2, 1, 5, 3, 4}
	// input2 := []int32{2, 5, 1, 3, 4}
	input := []int32{5, 1, 2, 3, 7, 8, 6, 4}
	input2 := []int32{1, 2, 5, 3, 4, 7, 8, 6}
	minimumBribes(input)
	minimumBribes(input2)

	fmt.Printf("sockMerch \n")
	//sockMerch
	arr := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	nPairs := sockMerchant(9, arr)
	fmt.Printf("%d \n", nPairs)

	// arrayManipulation
	fmt.Printf("arrayManipulation \n")
	rr1 := []int32{1, 5, 3}
	rr2 := []int32{4, 8, 7}
	rr3 := []int32{6, 9, 1}
	queries := [][]int32{rr1, rr2, rr3}

	res := arrayManipulation(10, queries)
	fmt.Printf("%d \n", res)

}
