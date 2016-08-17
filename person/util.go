package person

func randomBetween(min, max int) int {
	return randgen.Intn(max-min) + min
}

func splitDigits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	var digits []int
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	// Reverse the slice
	for i, j := len(digits)-1, 0; j < i; i, j = i-1, j+1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

func zeroLeftPad(digits []int, length int) []int {
	if len(digits) >= length {
		return digits
	}
	out := make([]int, length-len(digits), length)
	return append(out, digits...)
}
