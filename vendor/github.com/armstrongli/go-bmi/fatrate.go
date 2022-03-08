package gobmi

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	fatRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}

//package gobmi
//
//func FatRate(bmi float64, age int, sexWeight int) (fatRate float64) {
//	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
//	return
//}
