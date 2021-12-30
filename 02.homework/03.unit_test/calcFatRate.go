package unittest

import "fmt"

func calcFatRate(bmi float64, age int, sexWeight int) (fatRate float64, err error) {
	if bmi <= 0 {
		err = fmt.Errorf("BMI不能为0或者负数")
		return
	}
	if age <= 0 || age > 150 {
		err = fmt.Errorf("年龄不能小于等于0岁且大于150岁")
		return
	}
	if (sexWeight != 0) && (sexWeight != 1) {
		err = fmt.Errorf("性别权重只能为0或1")
		return
	}
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}
