package save

type SortsType map[string]bool

var defaultSorters = SortsType{ // Is Negative
	// Positives
	"Safety Index":                 false,
	"Climate Index":                false,
	"Health Care Index":            false,
	"Affordability Index":          false,
	"Quality of Life Index":        false,
	"Health Care Exp. Index":       false,
	"Purchasing Power Index":       false,
	"Local Purchasing Power Index": false,
	// ---------------------------------

	// Negatives
	"Rent Index":                                 true,
	"Crime Index":                                true,
	"Traffic Index":                              true,
	"Groceries Index":                            true,
	"Pollution Index":                            true,
	"Time Exp. Index":                            true,
	"CO2 Emission Index":                         true,
	"Inefficiency Index":                         true,
	"Exp Pollution Index":                        true,
	"Cost of Living Index":                       true,
	"Price To Income Ratio":                      true,
	"Restaurant Price Index":                     true,
	"Time Index (in minutes)":                    true,
	"Traffic Commute Time Index":                 true,
	"Cost of Living Plus Rent Index":             true,
	"Gross Rental Yield City Centre":             true,
	"Property Price to Income Ratio":             true,
	"Price To Rent Ratio City Centre":            true,
	"Mortgage As A Percentage Of Income":         true,
	"Gross Rental Yield Outside of Centre":       true,
	"Price To Rent Ratio Outside Of City Centre": true,
	// -------------------------------------------------
}

type SettingsType struct {
	Sorts SortsType
}

var Settings = SettingsType{defaultSorters}
