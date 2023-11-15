package parser

import (
	"github.com/nikdissv-forever/numbeogo/internal/consts"
	"github.com/nikdissv-forever/numbeogo/pkg/web"
	"golang.org/x/net/html"
	"strconv"
)

func ParseColumns(node *html.Node) []string {
	displayColumns := consts.ColumnOption.Select(node)[1:]
	columns := make([]string, len(displayColumns))
	for i, column := range displayColumns {
		columns[i] = web.GetText(column)
	}
	return columns
}

func GetName(node *html.Node) string {
	nameCol := consts.TableRowName.Select(node)
	if len(nameCol) == 0 {
		return ""
	}
	return web.GetText(nameCol[0])
}

func GetTableValues(node *html.Node) []float64 {
	columns := consts.TableColumns.Select(node)
	values := make([]float64, len(columns))
	for i, column := range columns {
		text := web.GetText(column)
		if val, err := strconv.ParseFloat(text, 32); err == nil {
			values[i] = val
		}
	}
	return values
}
