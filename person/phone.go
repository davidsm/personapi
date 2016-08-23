package person

import (
	"bytes"
	"strconv"
)

var mobilePrefixes = []string{
	"070",
	"072",
	"073",
	"076",
	"079",
}

func RandomPhoneNumber() string {
	// Only generates mobile numbers for now
	var buffer bytes.Buffer
	buffer.WriteString(mobilePrefixes[randgen.Intn(len(mobilePrefixes))])

	// Mobile phone numbers are always prefix + 7 digits
	for n := 0; n < 7; n++ {
		buffer.WriteString(strconv.Itoa(randgen.Intn(10)))
	}
	return buffer.String()
}
