package preferences

import (
	"github.com/ndenissov/numbeogo/gui/internal"
	"github.com/ndenissov/numbeogo/pkg"
)

var Preferences = internal.RootApp.Preferences()

func Prefixes(prefix string) func(string) string {
	return func(s string) string { return pkg.AddPrefix(prefix, s) }
}
