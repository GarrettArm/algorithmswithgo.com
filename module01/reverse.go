package module01

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	if len(word) == 0 {
		return ""
	}
	lastLetter := word[len(word)-1]
	firstPart := word[:len(word)-1]
	return fmt.Sprintf("%c%s", lastLetter, Reverse(firstPart))
}
