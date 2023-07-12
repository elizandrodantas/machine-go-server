package util

import "strconv"

func StringToInt(s string) (int, bool) {
	num, err := strconv.Atoi(s)

	if err != nil {
		return 0, false
	}

	return num, true
}
