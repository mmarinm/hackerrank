package main

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n+1)

	for _, q := range queries {
		a := q[0]
		b := q[1]
		k := int64(q[2])

		arr[a-1] += k
		arr[b] -= k
	}

	x := int64(0)
	max := int64(0)

	for _, el := range arr {
		x += el
		if x > max {
			max = x
		}
	}

	return max

}
