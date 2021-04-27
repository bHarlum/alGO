package module01

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"

func Reverse(word string) string {
	b := []rune(word)
	l := len(b)
	var r = make([]rune, len(b))
	for i, v := range b {
		if i > (l-1)/2 {
			fmt.Println("TRUE")
			return string(r)
		}
		r[l-1-i] = v
		r[i] = b[l-1-i]
	}

	return string(r)
}
