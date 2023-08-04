package numbeo

import (
	"github.com/nikdissv-forever/numbeogo/pkg/web"
	"strings"
)

type TableParams struct {
	URL, Title, Region string
}

func GetPageTable(params TableParams) Table {
	if params.Title != "" {
		params.Title = strings.Join([]string{"&title=", params.Title}, "")
	}
	if params.Region != "" {
		params.Region = strings.Join([]string{"&region=", params.Region}, "")
	}
	params.URL = strings.Join(
		[]string{params.URL, "?displayColumn=-1", params.Title, params.Region}, "")
	nodes, err := web.RequestNodes(params.URL)
	if err != nil {
		return Table{}
	}
	return ParseTable(nodes)
}
