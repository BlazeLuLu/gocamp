package main

import "fmt"

func QuickSort(arr *[]float64, start, end int) {
	// todo 确认终止条件，否则将无限递归下去

	pivotIdx := (start + end) / 2
	pivotV := (*arr)[pivotIdx]
	l, r := start, end
	for l <= r {
		for (*arr)[l] < pivotV {
			l++
		}
		for (*arr)[r] > pivotV {
			r--
		}
		if l >= r {
			break
		}

		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
	}
	// fmt.Println("l:", l, "r:", r)
	fmt.Println(*arr)
	if l == r {
		l++
		r--
	}
	if r > start {
		QuickSort(arr, start, r)
	}
	if l < end {
		QuickSort(arr, l, end)
	}
}
