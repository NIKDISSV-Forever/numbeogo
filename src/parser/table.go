package numbeo

import (
	"github.com/nikdissv-forever/numbeogo/internal/consts"
	"github.com/nikdissv-forever/numbeogo/internal/parser"
	"github.com/nikdissv-forever/numbeogo/recorder"
	"golang.org/x/net/html"
)

type Country struct {
	Name   string
	Values []float32
}

type Table struct {
	Headers       []string
	CountryValues []*Country
}

func ParseTable(node *html.Node) (parsed Table) {
	recorder.UpdateTitlesAndRegions(node)
	parsed.Headers = parser.ParseColumns(node)

	rowsNodes := consts.TableRows.Select(node)
	parsed.CountryValues = make([]*Country, len(rowsNodes))
	for i, row := range rowsNodes {
		name := parser.GetName(row)
		if name == "" {
			continue
		}
		parsed.CountryValues[i] = &Country{name, parser.GetTableValues(row)}
	}
	return
}
