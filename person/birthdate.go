package person

import (
	"time"
)

const maxAge int = 105

type BirthDate struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
}

func (bd *BirthDate) Age() int {
	now := time.Now()
	age := now.Year() - bd.Year
	if now.Month()-bd.Month < 0 {
		age--
	} else if now.Month()-bd.Month == 0 && now.Day() > bd.Day {
		age--
	}
	return age
}

func RandomBirthDate() BirthDate {
	now := time.Now()
	currentYear := now.Year()
	currentMonth := now.Month()
	currentDay := now.Day()

	// This rely on the auto-normalization of time.Date,
	// so that the date will always be valid (e.g. Apr 31 is converted to
	// May 1). This means that there will be a slight skew in the distribution,
	// as dates that can be "overflowed" to will appear somewhat more often.
	// For now, this behavior should be "good enough"
	year := currentYear - randgen.Intn(maxAge+1)

	availableMonths := 12
	if year == currentYear {
		availableMonths = int(currentMonth)
	}
	month := time.Month(randgen.Intn(availableMonths) + 1)

	availableDays := 31
	if year == currentYear && month == currentMonth {
		availableDays = currentDay
	}
	day := randgen.Intn(availableDays) + 1
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return BirthDate{Year: date.Year(), Month: date.Month(), Day: date.Day()}
}
