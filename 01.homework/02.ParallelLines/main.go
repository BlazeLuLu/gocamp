package main

import (
	"fmt"
	"math"
)

func main() {
	var lineSlopes []float64

	//分别输入两条直线两点的信息，并计算斜率
	for i := 0; i < 2; i++ {
		var pointOneCoordinates [2]float64
		fmt.Printf("输入第%d条直线的第一个点坐标x、y，请用空格分隔：", i+1)
		fmt.Scanln(&pointOneCoordinates[0], &pointOneCoordinates[1])

		var pointTwoCoordinates [2]float64
		fmt.Printf("输入第%d条直线的第二个点坐标x、y，请用空格分隔：", i+1)
		fmt.Scanln(&pointTwoCoordinates[0], &pointTwoCoordinates[1])

		lineSlope := math.Abs((pointOneCoordinates[1] - pointTwoCoordinates[1]) / (pointOneCoordinates[0] - pointTwoCoordinates[0]))
		//当两点x相同时，斜率为+Inf
		//fmt.Println(lineSlope)

		lineSlopes = append(lineSlopes, lineSlope)
	}

	//根据斜率判断直线是否平行
	if lineSlopes[0] == lineSlopes[1] {
		fmt.Println("直线1与直线2平行")
	} else {
		fmt.Println("直线1与直线2不平行")
	}

}
