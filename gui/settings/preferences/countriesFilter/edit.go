package countriesFilter

import (
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences"
)

const Prefix = "CountriesNames"

var AddPrefix = preferences.Prefixes(Prefix)

func Get(name string) bool {
	return preferences.GetBool(AddPrefix(name))
}

func Set(name string, showIt bool) {
	preferences.SetBool(AddPrefix(name), showIt)
}
