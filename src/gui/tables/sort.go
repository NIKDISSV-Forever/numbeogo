package tables

import (
	"github.com/nikdissv-forever/numbeogo/gui/settings/save"
	"github.com/nikdissv-forever/numbeogo/parser"
	"sort"
)

func sortFunction(headers []string, country *numbeo.Country) float32 {
	var pos, neg float32
	for i, value := range country.Values {
		if isNeg, ok := save.Settings.Sorts[headers[i]]; ok && isNeg {
			neg += value
		} else {
			pos += value
		}
	}
	return pos / neg
}

func sortTable(table numbeo.Table) { // Countries list
	sort.Slice(table.CountryValues, func(i, j int) bool {
		return sortFunction(table.Headers, table.CountryValues[i]) > sortFunction(table.Headers, table.CountryValues[j])
	})
}
