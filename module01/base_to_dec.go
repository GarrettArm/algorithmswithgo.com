package module01

import (
	"fmt"
	"math"
	"strconv"
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
	fmt.Printf("value: %v, base: %v\n", value, base)
	var output int
	valueStr := string(value)
	for i := 0; i < len(valueStr); i++ {
		x, _ := strconv.Atoi(string(valueStr[i]))
		reversedPosition := len(valueStr) - i - 1
		multiplier := math.Pow(float64(base), float64(reversedPosition))
		converted := float64(x) * multiplier
		fmt.Println(i, base, x, reversedPosition, multiplier, converted)
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
