package iteration

import "strings"

func Repeat(word string, count int) string {
	var repeated strings.Builder

	for i := 0; i < count; i++ {
		repeated.WriteString(word)
	}

	return repeated.String()
}
