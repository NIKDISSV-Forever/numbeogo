package sortRule

import "github.com/nikdissv-forever/numbeogo/gui/settings/preferences"

const Prefix = "sortsRule"

var AddPrefix = preferences.Prefixes(Prefix)

var defaultNegatives = [...]string{
	"Rent Index",
	"Crime Index",
	"Traffic Index",
	"Groceries Index",
	"Pollution Index",
	"Time Exp. Index",
	"CO2 Emission Index",
	"Inefficiency Index",
	"Exp Pollution Index",
	"Cost of Living Index",
	"Price To Income Ratio",
	"Restaurant Price Index",
	"Time Index (in minutes)",
	"Traffic Commute Time Index",
	"Cost of Living Plus Rent Index",
	"Gross Rental Yield City Centre",
	"Property Price to Income Ratio",
	"Price To Rent Ratio City Centre",
	"Mortgage As A Percentage Of Income",
	"Gross Rental Yield Outside of Centre",
	"Price To Rent Ratio Outside Of City Centre"}

func init() {
	if !get(Prefix) {
		for _, name := range defaultNegatives {
			Set(name, true)
		}
		set(Prefix, true)
	}
}
