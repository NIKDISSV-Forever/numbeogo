package tables

import (
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences/sortRule"
)

func SortFunction(headers []string, values []float64) float64 {
	var pos, neg float64 = 1, 1
	for i, value := range values {
		if sortRule.Get(headers[i]) {
			neg += value
		} else {
			pos += value
		}
	}
	return pos / neg
}
