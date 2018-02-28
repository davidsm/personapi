package person

import (
	"bytes"
	"strconv"
	"strings"
)

var domains = []string{
	"facebook.com",
	"gmail.com",
	"googlemail.com",
	"google.com",
	"hotmail.com",
	"mac.com",
	"mail.com",
	"msn.com",
	"live.com",
	"yahoo.com",
	"icloud.com",
	"inbox.com",
	"lavabit.com",
	"outlook.com",
	"protonmail.com",
	"telia.com",
	"bredband.net",
	"bahnhof.se",
	"tele2.se",
	"telenor.se",
	"comhem.se",
}

const (
	OptionFirstName        = iota
	OptionFirstAndLastName = iota
	OptionRandom           = iota
)

func createEmail(name string) string {
	domain := domains[randgen.Intn(len(domains))]
	var buffer bytes.Buffer
	buffer.WriteString(name)
	buffer.WriteString("@")
	buffer.WriteString(domain)
	return buffer.String()
}

func randomDigits(amount int) string {
	var buffer bytes.Buffer
	for i := 0; i < amount; i++ {
		buffer.WriteString(strconv.Itoa(randgen.Intn(10)))
	}
	return buffer.String()
}

func firstNameBased(firstName string) string {
	var buffer bytes.Buffer
	buffer.WriteString(firstName)
	numberLen := randomBetween(1, 6)
	buffer.WriteString(randomDigits(numberLen))
	return createEmail(buffer.String())
}

func firstAndLastNameBased(firstName, lastName string) string {
	var buffer bytes.Buffer

	separators := [3]string{".", "_", "-"}
	separator := separators[randgen.Intn(len(separators))]

	buffer.WriteString(firstName)
	buffer.WriteString(separator)
	buffer.WriteString(lastName)
	numberLen := randomBetween(1, 4)
	buffer.WriteString(randomDigits(numberLen))
	return createEmail(buffer.String())
}

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$&!?")

func randomAddress() string {
	var buffer bytes.Buffer
	length := randomBetween(8, 25)
	for i := 1; i < length; i++ {
		buffer.WriteRune(chars[randgen.Intn(len(chars))])
	}
	return createEmail(buffer.String())
}

func GenerateEmail(firstName, lastName string) string {
	replacer := strings.NewReplacer("å", "a", "ä", "a", "ö", "o")
	firstName = replacer.Replace(strings.ToLower(firstName))
	lastName = replacer.Replace(strings.ToLower(lastName))

	options := [3]int{OptionFirstName, OptionFirstAndLastName, OptionRandom}
	choice := options[randgen.Intn(len(options))]
	switch choice {
	case OptionFirstName:
		return firstNameBased(firstName)
	case OptionFirstAndLastName:
		return firstAndLastNameBased(firstName, lastName)
	case OptionRandom:
		return randomAddress()
	default:
		return randomAddress()
	}
}
