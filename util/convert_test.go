package util

import (
	"testing"
)

func TestConvertStringToNumber(t *testing.T) {
	numberTrue := "123"
	numberFalse := "a1b2c3"

	result1, testTrue := StringToInt(numberTrue)

	if !testTrue {
		t.Errorf("expected `true` and received `%t`", testTrue)
	}

	if result1 != 123 {
		t.Error("expected number does not match")
	}

	_, testFalse := StringToInt(numberFalse)

	if testFalse {
		t.Errorf("expected `false` and received `%t`", testFalse)
	}

}
