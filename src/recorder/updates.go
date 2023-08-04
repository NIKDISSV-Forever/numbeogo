package recorder

import (
	"github.com/nikdissv-forever/numbeogo/internal/consts"
	"github.com/nikdissv-forever/numbeogo/internal/mutex"
	"github.com/nikdissv-forever/numbeogo/pkg/web"
	"golang.org/x/net/html"
	"net/url"
)

func UpdateTitles(node *html.Node) {
	if !record {
		return
	}
	mutex.Locker.Lock()
	defer mutex.Locker.Unlock()
	for _, option := range consts.TitleOption.Select(node) {
		key := web.GetText(option)
		if _, exist := Recorded.Titles[key]; !exist {
			if value := web.GetAttr(option, "value"); value != "" {
				Recorded.Titles[key] = value
			}
		}
	}
	Signal.Bell()
}

func UpdateRegions(node *html.Node) {
	if !record {
		return
	}
	mutex.Locker.Lock()
	defer mutex.Locker.Unlock()
	for _, link := range consts.RegionLinks.Select(node) {
		key := web.GetText(link)
		if _, exist := Recorded.Regions[key]; !exist {
			if href := web.GetAttr(link, "href"); href != "" {
				if parse, err := url.Parse(href); err == nil {
					Recorded.Regions[key] = parse.Query().Get("region")
				}
			}
		}
	}
	Signal.Bell()
}

func UpdateTitlesAndRegions(node *html.Node) {
	UpdateTitles(node)
	UpdateRegions(node)
}
