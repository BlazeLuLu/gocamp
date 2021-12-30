package unittest

import "testing"

func Test_CalcBMI_1(t *testing.T) {
	_, err := calcBMI(-1, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_CalcBMI_2(t *testing.T) {
	_, err := calcBMI(0, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_CalcBMI_3(t *testing.T) {
	_, err := calcBMI(1, -1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_CalcBMI_4(t *testing.T) {
	_, err := calcBMI(1, 0)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_CalcBMI_5(t *testing.T) {
	var expectedBMI float64 = 1
	bmi, err := calcBMI(1, 1)
	if err != nil {
		t.Errorf("should be error, but err returned is nil")
	}
	if bmi != expectedBMI {
		t.Error("expected is different from calculated")
	}
}
