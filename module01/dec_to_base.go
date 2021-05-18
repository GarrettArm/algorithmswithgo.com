package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//

func DecToBase(dec, base int) string {
	letterNumbers := map[string]string{
		"10": "A",
		"11": "B",
		"12": "C",
		"13": "D",
		"14": "E",
		"15": "F",
	}

	var result string
	for dec != 0 {
		remainder := fmt.Sprint(dec % base)
		if letter, ok := letterNumbers[remainder]; ok {
			remainder = letter
		}
		dec = dec / base
		result = fmt.Sprint(remainder) + result
	}
	return result
}
