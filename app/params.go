package app

import (
	"net/url"
	"strconv"
)

const (
	paramAmount = "amount"
)

const (
	minAmount     = 1
	maxAmount     = 100
	defaultAmount = 10
)

func readAmount(params url.Values) int {
	val := params.Get(paramAmount)
	if len(val) == 0 {
		return defaultAmount
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return defaultAmount
	}
	if num < minAmount {
		return minAmount
	} else if num > maxAmount {
		return maxAmount
	}
	return num
}
