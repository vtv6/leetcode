package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	a := arr{boxTypes}
	sort.Sort(a)
	result, index := 0, 0
	for truckSize > 0 && index < len(boxTypes) {
		if truckSize <= a.a[index][0] {
			result += truckSize * a.a[index][1]
			truckSize = 0
		} else {
			result += a.a[index][0] * a.a[index][1]
			truckSize -= a.a[index][0]
		}
		index++
	}

	return result
}

type arr struct {
	a [][]int
}

func (a arr) Len() int {
	return len(a.a)
}

func (a arr) Less(i, j int) bool {
	return a.a[i][1] > a.a[j][1]
}

func (a arr) Swap(i, j int) {
	a.a[i], a.a[j] = a.a[j], a.a[i]
}
