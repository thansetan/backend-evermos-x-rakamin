package utils

import "strconv"

func StringToUint(str string) (uint, error) {
	res, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(res), nil
}
