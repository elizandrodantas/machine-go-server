package util

import (
	"testing"
)

func TestValidatorBetween(t *testing.T) {
	valueInvalid := ValidatorParamWithBetween("a", "test", 2, 4)
	valueInvalid2 := ValidatorParamWithBetween("abcde", "test", 2, 4)

	if valueInvalid == nil {
		t.Error("expected an error [1]")
	}

	if valueInvalid2 == nil {
		t.Error("expected an error [2]")
	}

	valueValid := ValidatorParamWithBetween("abc", "test", 1, 5)

	if valueValid != nil {
		t.Errorf("correct parameter should not return error, and returned: %s", valueValid)
	}
}
