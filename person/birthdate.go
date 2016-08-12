package person

import "time"

type BirthDate struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
}

func (bd *BirthDate) age(now time.Time) int {
	age := now.Year() - bd.Year

	currentMonth := now.Month()
	if currentMonth-bd.Month < 0 {
		age--
	} else if currentMonth-bd.Month == 0 && now.Day() < bd.Day {
		age--
	}
	return age
}

func (bd *BirthDate) Age() int {
	now := time.Now()
	return bd.age(now)
}

func randomBetween(min, max int) int {
	return randgen.Intn(max-min) + min
}

func RandomBirthDate(minAge, maxAge int) *BirthDate {
	now := time.Now()
	currentYear, currentMonth, currentDay := now.Year(), now.Month(), now.Day()

	// This rely on the auto-normalization of time.Date,
	// so that the date will always be valid (e.g. Apr 31 is converted to
	// May 1). This means that there will be a slight skew in the distribution,
	// as dates that can be "overflowed" to will appear somewhat more often.
	// For now, this behavior should be "good enough"
	minYear := currentYear - (maxAge + 1)
	maxYear := currentYear - minAge
	year := randomBetween(minYear, maxYear+1)

	minMonth := time.January
	maxMonth := time.December
	if year == currentYear || year == maxYear {
		maxMonth = currentMonth
	} else if year == minYear {
		minMonth = currentMonth
	}
	month := time.Month(randomBetween(int(minMonth), int(maxMonth)+1))

	minDay := 1
	maxDay := 31
	if year == currentYear && month == currentMonth {
		maxDay = currentDay
	} else if year == maxYear && month == maxMonth {
		maxDay = currentDay
	} else if year == minYear && month == minMonth {
		minDay = currentDay + 1
	}

	day := randomBetween(minDay, maxDay+1)

	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return &BirthDate{Year: date.Year(), Month: date.Month(), Day: date.Day()}
}
