package guiNumbeo

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences/countriesFilter"
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences/sortRule"
	"github.com/nikdissv-forever/numbeogo/gui/tables"
	"github.com/nikdissv-forever/numbeogo/parser"
	"golang.org/x/image/colornames"
	"sort"
)

const columns = 4

type countryData struct {
	Overall      float64
	InCategories map[string][]float64
}
type countryValue struct {
	Name       string
	Overall    float64
	Categories []string
}

type ContentSetter struct {
	Data   map[string]*numbeo.Table
	show   []*countryValue
	widget *fyne.Container
	spacer *layout.Spacer
}

func NewContentSetter(data map[string]*numbeo.Table) *ContentSetter {
	return &ContentSetter{Data: data, show: make([]*countryValue, 0), spacer: &layout.Spacer{}}
}

func (c *ContentSetter) sort(i int, j int) bool {
	return c.show[i].Overall >= c.show[j].Overall
}

func (c *ContentSetter) Resort() {
	sort.Slice(c.show, c.sort)
}

func (c *ContentSetter) GetWidget() *container.Scroll {
	if c.widget == nil {
		c.widget = container.NewGridWithColumns(columns)
	}
	return container.NewScroll(c.widget)
}

func (c *ContentSetter) Refresh() {
	c.UpdateData()
	c.Resort()
	c.UpdateWidgetContent()
}

func (c *ContentSetter) requiredSizeCategory(headers []string) int {
	return (len(headers)-1)*4 + 2
}

func (c *ContentSetter) UpdateData() {
	show := make(map[string]*countryData)
	categories := make([]string, 0, len(c.Data))
	for cat, table := range c.Data {
		for name, values := range table.CountryValues {
			country, ok := show[name]
			if !ok {
				country = &countryData{InCategories: make(map[string][]float64, len(c.Data))}
				show[name] = country
			}
			country.InCategories[cat] = values
			country.Overall += tables.SortFunction(table.Headers, values)
		}
		categories = append(categories, cat)
	}
	sort.Strings(categories)
	c.show = make([]*countryValue, 0, len(c.show))
	for name, data := range show {
		if !countriesFilter.Get(name) {
			continue
		}
		country := &countryValue{
			Name:       name,
			Overall:    data.Overall,
			Categories: make([]string, 0, len(categories)),
		}
		c.show = append(c.show, country)
		for _, cat := range categories {
			country.Categories = append(country.Categories, cat)
		}
	}
}
func (c *ContentSetter) requiredSize() int {
	sz := len(c.show)
	for _, country := range c.show {
		var hi int
		var category string
		var values []float64
		name := country.Name
		for hi, category = range country.Categories {
			if values = c.getValues(category, name); len(values) != 0 {
				break
			}
		}
		sz += c.requiredSizeCategory(c.Data[country.Categories[hi]].Headers) + 1
		for _, category = range country.Categories[hi+1:] {
			if values = c.getValues(category, name); len(values) == 0 {
				continue
			}
			sz += c.requiredSizeCategory(c.Data[category].Headers) + 2
		}
	}
	return sz
}

func (c *ContentSetter) makeCategory(headers []string, values []float64) {
	c.widget.Objects = append(c.widget.Objects,
		c.newHeader(headers[0]), c.newValue(values[0]))
	for hi := 1; hi < len(headers); hi++ {
		c.widget.Objects = append(c.widget.Objects,
			c.spacer, c.spacer,
			c.newHeader(headers[hi]), c.newValue(values[hi]))
	}
}

func (c *ContentSetter) UpdateWidgetContent() {
	c.GetWidget()
	if sz := c.requiredSize(); cap(c.widget.Objects) < sz {
		c.widget.Objects = make([]fyne.CanvasObject, 0, sz)
	}
	c.widget.Objects = c.widget.Objects[:0]

	for i, country := range c.show {
		name := country.Name
		c.widget.Objects = append(c.widget.Objects, c.newTitle(i+1, name))

		var hi int
		var category string
		var values []float64
		for hi, category = range country.Categories {
			if values = c.getValues(category, name); len(values) != 0 {
				break
			}
		}

		category = country.Categories[hi]
		cat := c.newCategory(category)
		cat.TextStyle.Italic = true
		c.widget.Objects = append(c.widget.Objects, cat)
		c.makeCategory(c.Data[category].Headers, values)
		for _, category = range country.Categories[hi+1:] {
			values = c.getValues(category, name)
			if len(values) == 0 {
				continue
			}
			c.widget.Objects = append(c.widget.Objects, c.spacer, c.newCategory(category))
			c.makeCategory(c.Data[category].Headers, values)
		}
	}
	println(len(c.widget.Objects), cap(c.widget.Objects))
	c.widget.Refresh()
}

func (c *ContentSetter) getValues(category, name string) []float64 {
	return c.Data[category].CountryValues[name]
}
func (c *ContentSetter) newCategory(category string) *canvas.Text {
	text := canvas.NewText(category, colornames.Burlywood)
	text.TextStyle.Bold = true
	return text
}

func (c *ContentSetter) newHeader(header string) fyne.CanvasObject {
	var button *widget.Button
	button = widget.NewButton(header, func() {
		button.Importance = c.buttonImportance(sortRule.Switch(header))
		c.Refresh()
	})
	button.Importance = c.buttonImportance(sortRule.Get(header))
	return button
}
func (c *ContentSetter) newValue(value float64) *canvas.Text {
	text := canvas.NewText(fmt.Sprintf("%.1f", value), colornames.Green)
	text.TextStyle.Italic = true
	return text
}

func (c *ContentSetter) newTitle(rating int, name string) *canvas.Text {
	text := canvas.NewText(fmt.Sprintf("%d\t%s", rating, name), colornames.Goldenrod)
	text.TextStyle.Italic = true
	text.TextStyle.Bold = true
	return text
}

func (c *ContentSetter) buttonImportance(isNeg bool) widget.ButtonImportance {
	if isNeg {
		return widget.LowImportance
	}
	return widget.HighImportance
}
