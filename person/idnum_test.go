package person

import (
	"reflect"
	"testing"
	"time"
)

var splitTests = []struct {
	in       int
	expected []int
}{
	{0, []int{0}},
	{5, []int{5}},
	{98, []int{9, 8}},
	{123, []int{1, 2, 3}},
	{100, []int{1, 0, 0}},
}

func TestSplitDigits(t *testing.T) {
	for _, test := range splitTests {
		digits := splitDigits(test.in)
		if !reflect.DeepEqual(digits, test.expected) {
			t.Errorf("splitDigits(%d): Got %v, expected %v", test.in, digits, test.expected)
		}
	}
}

var leftPadTests = []struct {
	in       []int
	length   int
	expected []int
}{
	{[]int{2}, 1, []int{2}},
	{[]int{2}, 2, []int{0, 2}},
	{[]int{2, 3}, 2, []int{2, 3}},
	{[]int{2, 3}, 4, []int{0, 0, 2, 3}},
	{[]int{}, 2, []int{0, 0}},
}

func TestZeroLeftPad(t *testing.T) {
	for _, test := range leftPadTests {
		digits := zeroLeftPad(test.in, test.length)
		if !reflect.DeepEqual(digits, test.expected) {
			t.Errorf("zeroLeftPad(%v, %d): Got %v, expected %v", test.in, test.length, digits, test.expected)
		}
	}
}

var controlNumberTests = []struct {
	in       []int
	expected int
}{
	{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2},
	{[]int{6, 4, 0, 8, 2, 3, 3, 2, 3}, 4},
	{[]int{5, 5, 0, 5, 1, 2, 2, 4, 0}, 7},
	{[]int{2, 5, 1, 1, 1, 5, 4, 8, 7}, 0},
}

func TestCalculateControlNumber(t *testing.T) {
	for _, test := range controlNumberTests {
		controlNumber := calculateControlNumber(test.in)
		if controlNumber != test.expected {
			t.Errorf("calculateControlNumber(%v): Got %d, expected %d", test.in,
				controlNumber, test.expected)
		}
	}
}

func TestGenerateIdNumber(t *testing.T) {
	// Bit awkward to test things which use random generation.
	// TODO: Look into setting seed to a fixed value for tests
	bd := BirthDate{1925, time.November, 15}
	idNum := GenerateIdNumber(bd, GenderMale)
	t.Log("IdNum:", idNum)
	if len(idNum) != 11 {
		t.Fatal("Expected generated id number to be 11 characters, was", len(idNum))
	}
	if string(idNum[6]) != "-" {
		t.Error("Expected separator to be \"-\" was", string(idNum[6]))
	}

	bd = BirthDate{1915, time.November, 15}
	idNum = GenerateIdNumber(bd, GenderMale)
	t.Log("IdNum:", idNum)
	if len(idNum) != 11 {
		t.Fatal("Expected generated id number to be 11 characters, was", len(idNum))
	}
	if string(idNum[6]) != "+" {
		t.Error("Expected separator to be \"+\" was", string(idNum[6]))
	}
}
