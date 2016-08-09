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

// calculateControlNumber calculates the control digit of a Swedish
// personal identification number ("personnummer")
// Algorithm can be found at https://www.skatteverket.se/privat/folkbokforing/personnummer/personnumretsuppbyggnad
func calculateControlNumber(digits []int) int {
	sum := 0
	// Multiply each number with 2 or 1 alternatively
	for i, digit := range digits {
		var multiple int
		if i%2 == 0 {
			multiple = 2
		} else {
			multiple = 1
		}

		// Add each result to a total sum. Split numbers >= 10 into individual digits
		for _, digit := range splitDigits(digit * multiple) {
			sum += digit
		}
	}

	// The control number is 10 minus the last digit of the previously
	// calculated sum. If it's calculated to 10, 0 is used instead
	controlNumber := 10 - (sum % 10)
	if controlNumber == 10 {
		return 0
	} else {
		return controlNumber
	}
}
