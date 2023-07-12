package util

import (
	"fmt"
)

func ValidatorParamWithBetween(data, paramName string, min, max int) error {
	if len(data) < min || len(data) > max {
		if len(data) < min {
			return fmt.Errorf("%s must be more than %d characters", paramName, min)
		}

		return fmt.Errorf("%s must be less than %d characters", paramName, max)
	}

	return nil
}
