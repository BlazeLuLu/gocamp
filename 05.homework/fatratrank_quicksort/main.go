package main

import (
	"fmt"
	"log"
)

func main() {
	input := &inputFromStd{}
	calc := &Calc{}
	rk := &FatRateRank{}
	records := NewRecord("/Users/luyue/Downloads/records.txt")

	for {
		pi := input.GetInput()
		if err := records.savePersonalInformation(pi); err != nil {
			log.Fatal("保存失败：", err)
		}
		fr, err := calc.FatRate(pi)
		if err != nil {
			log.Fatal("计算体脂率失败：", err)
		}
		rk.inputRecord(pi.Name, fr)

		rankResult, _ := rk.getRank(pi.Name)
		fmt.Println("排名结果：", rankResult)
	}
}
