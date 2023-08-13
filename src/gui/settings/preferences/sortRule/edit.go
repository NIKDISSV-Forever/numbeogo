package sortRule

import (
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences"
	"strings"
)

func AddPrefix(name string) string {
	return strings.Join([]string{Prefix, name}, "")
}

var (
	get = preferences.Preferences.Bool
	set = preferences.Preferences.SetBool
)

func Get(name string) bool {
	return get(AddPrefix(name))
}

func Set(name string, isNegative bool) {
	set(AddPrefix(name), isNegative)
}

func Switch(name string) bool {
	name = AddPrefix(name)
	newVal := !get(name)
	set(name, newVal)
	return newVal
}
