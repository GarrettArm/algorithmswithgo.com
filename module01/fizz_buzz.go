package module01

import (
	"fmt"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	output := make([]string, n)
	for i := 0; i < n+1; i++ {
		if i == 0 {
			continue
		}
		switch {
		case i%15 == 0:
			output[i-1] = "Fizz Buzz"
		case i%3 == 0:
			output[i-1] = "Fizz"
		case i%5 == 0:
			output[i-1] = "Buzz"
		default:
			output[i-1] = fmt.Sprint(i)
		}
	}
	fmt.Println(strings.Join(output, ", "))
}
