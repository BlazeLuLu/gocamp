package unittest

import "fmt"

func calcBMI(weight float64, tall float64) (bmi float64, err error) {
	if weight <= 0 {
		err = fmt.Errorf("体重不能为0或者负数")
		return
	}
	if tall <= 0 {
		err = fmt.Errorf("身高不能为0或者负数")
		return
	}
	bmi = weight / (tall * tall)
	return
}
