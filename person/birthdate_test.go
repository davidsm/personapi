package person

import (
	"testing"
	"time"
)

var ageTests = []struct {
	bd       BirthDate
	expected int
}{
	{BirthDate{1986, time.Month(5), 16}, 30},
	{BirthDate{2000, time.Month(10), 8}, 15},
	{BirthDate{1950, time.Month(8), 25}, 65},
	{BirthDate{2016, time.Month(8), 8}, 0},
}

func TestAge(t *testing.T) {
	now := time.Date(2016, 8, 10, 0, 0, 0, 0, time.UTC)
	for _, test := range ageTests {
		age := test.bd.age(now)
		if age != test.expected {
			t.Errorf("bd.Age() (bd=%v): Got %d, expected %d", test.bd, age, test.expected)
		}
	}
}
