package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func StringToUint(str string) (uint, error) {
	if str == "" {
		return 0, nil
	}
	res, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(res), nil
}

func GenerateProductSlug(str string) string {
	return strings.ToLower(strings.Join(strings.Split(str, " "), "-"))
}

func StringToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	res, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func GenerateInvoiceNumber() string {
	return fmt.Sprintf("INV-%d", time.Now().Unix())
}
