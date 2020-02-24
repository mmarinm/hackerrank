package main

import "fmt"

func minimumBribes(q []int32) {
	n := int32(len(q))
	cnt := int32(0)
	for i := n - 1; i >= 0; i-- {
		if i+1 != q[i] {
			if i-1 >= 0 && i+1 == q[i-1] {
				cnt++
				q[i], q[i-1] = q[i-1], q[i]
				continue
			}

			if i-2 >= 0 && i+1 == q[i-2] {
				cnt += 2
				q[i-2], q[i-1] = q[i-1], q[i-2]
				q[i-1], q[i] = q[i], q[i-1]
				continue
			}
			fmt.Println("Too chaotic")
			return

		}
	}
	fmt.Println(cnt)
	return
}
