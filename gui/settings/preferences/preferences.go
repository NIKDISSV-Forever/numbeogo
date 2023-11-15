package preferences

import (
	"github.com/nikdissv-forever/numbeogo/gui/internal"
	"github.com/nikdissv-forever/numbeogo/pkg"
)

var Preferences = internal.RootApp.Preferences()

func Prefixes(prefix string) func(string) string {
	return func(s string) string { return pkg.AddPrefix(prefix, s) }
}
