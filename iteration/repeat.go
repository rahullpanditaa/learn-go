package iteration

import "strings"

func Repeat(character string, reps int) string {
	if reps == 0 {
		return ""
	}

	var repeated strings.Builder
	for range reps {
		repeated.WriteString(character)
	}
	return repeated.String()
}
