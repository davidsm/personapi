package person

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

func calculateControlNumber(digits []int) int {
	// Multiply each number with 2 or 1 alternatively
	multiplied := make([]int, len(digits))
	for i, digit := range digits {
		var multiple int
		if i%2 == 0 {
			multiple = 2
		} else {
			multiple = 1
		}
		multiplied[i] = digit * multiple
	}

	// Calculate the sum. Split numbers >= 10 into individual digits
	sum := 0
	for _, num := range multiplied {
		digits := splitDigits(num)
		for _, digit := range digits {
			sum += digit
		}
	}

	// The control number is 10 minus the last digit of the previously
	// calculated sum
	return 10 - (sum % 10)
}
