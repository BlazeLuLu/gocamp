package unittest

import "testing"

func Test_calcFatRate_1(t *testing.T) {
	_, err := calcFatRate(-1, 1, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_calcFatRate_2(t *testing.T) {
	_, err := calcFatRate(0, 1, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_calcFatRate_3(t *testing.T) {
	_, err := calcFatRate(1, -1, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_calcFatRate_4(t *testing.T) {
	_, err := calcFatRate(1, 0, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_calcFatRate_6(t *testing.T) {
	_, err := calcFatRate(1, 151, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_calcFatRate_7(t *testing.T) {
	_, err := calcFatRate(1, 1, 3)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func Test_calcFatRate_8(t *testing.T) {
	var expectedFateRate float64 = 0.241
	FatRate, err := calcFatRate(26.30, 38, 1)
	if err != nil {
		t.Errorf("should be error, but err returned is nil")
	}
	if FatRate != expectedFateRate {
		t.Error("expected is different from calculated")
	}
}
