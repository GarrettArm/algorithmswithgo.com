package module01

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var rev string
	for i := len(word); i > 0; i-- {
		rev = fmt.Sprintf("%s%c", rev, word[i-1])
	}
	return rev

}
