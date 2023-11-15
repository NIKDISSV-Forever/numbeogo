package recorder

import (
	"github.com/nikdissv-forever/numbeogo/internal/consts"
	"github.com/nikdissv-forever/numbeogo/internal/mutex"
	"github.com/nikdissv-forever/numbeogo/pkg/web"
	"golang.org/x/net/html"
	"net/url"
)

func (recorder *Container) UpdateTitles(node *html.Node) {
	mutex.Locker.Lock()
	defer mutex.Locker.Unlock()
	if !record {
		return
	}
	for _, option := range consts.TitleOption.Select(node) {
		key := web.GetText(option)
		if _, exist := recorder.Titles[key]; !exist {
			if value := web.GetAttr(option, "value"); value != "" {
				recorder.Titles[key] = value
			}
		}
	}
	Signal.Bell()
}

func (recorder *Container) UpdateRegions(node *html.Node) {
	mutex.Locker.Lock()
	defer mutex.Locker.Unlock()
	if !record {
		return
	}
	for _, link := range consts.RegionLinks.Select(node) {
		key := web.GetText(link)
		if _, exist := recorder.Regions[key]; !exist {
			if href := web.GetAttr(link, "href"); href != "" {
				if parse, err := url.Parse(href); err == nil {
					recorder.Regions[key] = parse.Query().Get("region")
				}
			}
		}
	}
	Signal.Bell()
}
func (recorder *Container) AddCountry(name string) {
	mutex.Locker.Lock()
	defer mutex.Locker.Unlock()
	if !record {
		return
	}
	recorder.Countries.Insert(name)
}

func (recorder *Container) UpdateTitlesAndRegions(node *html.Node) {
	recorder.UpdateTitles(node)
	recorder.UpdateRegions(node)
}
