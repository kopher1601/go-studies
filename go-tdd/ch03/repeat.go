package ch03

import "strings"

func Repeat(character string) string {
	builder := strings.Builder{}
	for i := 0; i < 5; i++ {
		builder.WriteString(character)
	}
	return builder.String()
}
