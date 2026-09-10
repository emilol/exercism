package raindrops

import "strconv"

func Convert(number int) string {
	var value = ""
	if number%3 == 0 {
		value = "Pling"
	}
	if number%5 == 0 {
		value = value + "Plang"
	}
	if number%7 == 0 {
		value = value + "Plong"
	}

	if value != "" {
		return value
	}

	return strconv.Itoa(number)
}
