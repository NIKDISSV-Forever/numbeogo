package internal

import (
	"regexp"
	"strings"
)

func joinQuery(query string) string {
	return strings.Join([]string{"(?i).*", query, ".*"}, "")
}

func GetFilterFunc(query string) func(string) bool {
	var err error
	var compile *regexp.Regexp
	if compile, err = regexp.Compile(joinQuery(query)); err != nil {
		compile = regexp.MustCompile(joinQuery(regexp.QuoteMeta(query)))
	}
	return compile.MatchString
}
