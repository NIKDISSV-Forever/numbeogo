package settings

import (
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences"
)

const Prefix = "Settings"

var AddPrefix = preferences.Prefixes(Prefix)

func GetInt(name string) int {
	return preferences.Preferences.Int(AddPrefix(name))
}

func SetInt(name string, value int) {
	preferences.Preferences.SetInt(AddPrefix(name), value)
}

func SetPgSz(value int) {
	SetInt("PgSz", value)
}

func GetPgSz() int {
	return GetInt("PgSz")
}

func init() {
	if GetPgSz() == 0 {
		SetPgSz(3)
	}
}
