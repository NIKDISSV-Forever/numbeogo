package numbeo

import (
	"github.com/nikdissv-forever/numbeogo/internal/consts"
	"github.com/nikdissv-forever/numbeogo/internal/parser"
	"github.com/nikdissv-forever/numbeogo/recorder"
	"golang.org/x/net/html"
)

type Country struct {
	Name   string
	Values []float64
}

type Table struct {
	Headers       []string
	CountryValues map[string][]float64
}

func ParseTable(node *html.Node) *Table {
	recorder.Recorded.UpdateTitlesAndRegions(node)
	parsed := &Table{}
	parsed.Headers = parser.ParseColumns(node)
	rowsNodes := consts.TableRows.Select(node)
	parsed.CountryValues = make(map[string][]float64, len(rowsNodes))
	for _, row := range rowsNodes {
		name := parser.GetName(row)
		if name == "" {
			continue
		}
		recorder.Recorded.AddCountry(name)
		parsed.CountryValues[name] = parser.GetTableValues(row)
	}
	return parsed
}
