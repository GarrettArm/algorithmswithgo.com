package module01

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
	var result string
	charset := "0123456789ABCDEF"
	for dec != 0 {
		remainder := dec % base
		result = string(charset[remainder]) + result
		dec = dec / base
	}
	return result
}
