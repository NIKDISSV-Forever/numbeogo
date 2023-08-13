package countriesFilter

import (
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences"
	"strings"
)

const NamesPrefix = "CountriesNames"

func AddPrefix(name string) string {
	return strings.Join([]string{NamesPrefix, name}, "")
}

func Get(name string) bool {
	return preferences.GetBool(AddPrefix(name))
}

func Set(name string, showIt bool) {
	preferences.SetBool(AddPrefix(name), showIt)
}
