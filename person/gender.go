package person

type Gender string

const (
	GenderFemale Gender = "female"
	GenderMale   Gender = "male"
)

func RandomGender() Gender {
	genders := [2]Gender{GenderFemale, GenderMale}
	return genders[randgen.Intn(len(genders))]
}
