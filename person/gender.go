package person

type gender string

const (
	GenderFemale gender = "female"
	GenderMale   gender = "male"
)

func RandomGender() gender {
	genders := [2]gender{GenderFemale, GenderMale}
	return genders[randgen.Intn(len(genders))]
}
