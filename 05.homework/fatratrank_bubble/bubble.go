package main

import "fmt"

func Bubble(arr *[]float64) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
		// fmt.Println("中间状态：", *arr)
	}
	fmt.Println("最终状态：", *arr)
}
