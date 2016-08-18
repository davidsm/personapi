package person

import (
	"bytes"
	"strconv"
)

// calculateControlNumber calculates the control digit of a Swedish
// personal identification number ("personnummer")
// Algorithm can be found at https://www.skatteverket.se/privat/folkbokforing/personnummer/personnumretsuppbyggnad
func calculateControlNumber(digits []int) int {
	sum := 0
	// Multiply each number with 2 or 1 alternately
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
	return (10 - (sum % 10)) % 10
}

// generateBirthNumber generates the "birth number".
// The number should be a number between 1 and 999,
// The number is odd for men and even for women
func generateBirthNumber(gender Gender) int {
	// Number 1-998
	num := randgen.Intn(998) + 1
	isEven := num%2 == 0
	if (gender == GenderMale && isEven) || (gender == GenderFemale && !isEven) {
		num++
	}
	return num
}

func formatIdNumber(digits []int, separator string) string {
	var buffer bytes.Buffer
	for _, digit := range digits[:6] {
		buffer.WriteString(strconv.Itoa(digit))
	}
	buffer.WriteString(separator)
	for _, digit := range digits[6:] {
		buffer.WriteString(strconv.Itoa(digit))
	}
	return buffer.String()
}

func GenerateIdNumber(birthDate *BirthDate, gender Gender) string {
	digits := make([]int, 0, 9)
	yearDigits := splitDigits(birthDate.Year)
	decadeAndYearDigits := yearDigits[len(yearDigits)-2:]
	digits = append(digits, decadeAndYearDigits...)
	digits = append(digits, zeroLeftPad(splitDigits(int(birthDate.Month)), 2)...)
	digits = append(digits, zeroLeftPad(splitDigits(birthDate.Day), 2)...)
	digits = append(digits, zeroLeftPad(splitDigits(generateBirthNumber(gender)), 3)...)
	controlNumber := calculateControlNumber(digits)
	digits = append(digits, controlNumber)
	var separator string
	if birthDate.Age() >= 100 {
		separator = "+"
	} else {
		separator = "-"
	}
	return formatIdNumber(digits, separator)
}
