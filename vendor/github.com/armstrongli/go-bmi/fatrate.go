package gobmi

func FatRate(bmi float64, age int, sexWeight int) (fatRate float64) {
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}
