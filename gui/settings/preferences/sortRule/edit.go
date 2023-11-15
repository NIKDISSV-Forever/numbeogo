package sortRule

import (
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences"
)

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
