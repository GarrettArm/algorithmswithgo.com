package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var output int
	charset := "0123456789ABCDEF"
	for i := 0; i < len(value); i++ {
		x := strings.Index(charset, string(value[i]))
		reversedPosition := len(value) - i - 1
		multiplier := math.Pow(float64(base), float64(reversedPosition))
		converted := float64(x) * multiplier
		output += int(converted)
	}
	return output
}

/*
	1  	2	1
	10 	2   	2
	21	3	7
	1110	2	14
*/
