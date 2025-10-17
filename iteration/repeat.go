package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var result strings.Builder
	for i := 0; i < repeatCount; i++ {
		result.WriteString(character)
	}
	return result.String()
}
