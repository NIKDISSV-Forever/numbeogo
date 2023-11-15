package pkg

import "strings"

func AddPrefix(prefix, value string) string {
	var b strings.Builder
	b.Grow(len(prefix) + len(value))
	b.WriteString(prefix)
	b.WriteString(value)
	return b.String()
}
