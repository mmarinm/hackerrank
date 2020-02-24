package main

func sockMerchant(n int32, ar []int32) int32 {
	if n <= 1 {
		return 0
	}

	m := make(map[int]int)

	for i := 0; i < int(n); i++ {
		_, ok := m[int(ar[i])]
		if ok {
			m[int(ar[i])]++
		} else {
			m[int(ar[i])] = 1
		}
	}

	nPairs := 0
	for _, value := range m {
		nPairs += (value - value%2) / 2
	}

	return int32(nPairs)

}
