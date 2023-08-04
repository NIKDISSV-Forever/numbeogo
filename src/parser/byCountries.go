package numbeo

import (
	"github.com/nikdissv-forever/numbeogo/internal/consts"
	"github.com/nikdissv-forever/numbeogo/pkg/web"
)

type Category struct {
	Name, URL string
}

// GetRankingByCountryCategories
// will return a map
// with keys as category names (pass to TableParams.Category)
// and values as links to the numbeo.com page (pass to TableParams.URL)
func GetRankingByCountryCategories() ([]*Category, error) {
	content, err := web.RequestNodes(consts.IndexPage)
	if err != nil {
		return []*Category{}, err
	}
	nodes := consts.ByCountryLinks.Select(content)
	categories := make([]*Category, 0, len(nodes))
	for _, node := range nodes {
		href := web.GetAttr(node, "href")
		if href == "" {
			continue
		}
		name := consts.RemoveSuffix.ReplaceAllString(web.GetText(node), "")
		categories = append(categories, &Category{name, href})
	}
	return categories, nil
}
