package guiNumbeo

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/nikdissv-forever/numbeogo/gui/settings"
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences/countriesFilter"
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences/sortRule"
	"github.com/nikdissv-forever/numbeogo/gui/tables"
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

type limits struct {
	start, end int
}

type ContentSetter struct {
	Data   map[string]*numbeo.Table
	show   []*countryValue
	widget *fyne.Container
	spacer *layout.Spacer

	scroll *container.Scroll
	data   []fyne.CanvasObject
	parts  []limits
	curs   int
}

func NewContentSetter(data map[string]*numbeo.Table) *ContentSetter {
	return &ContentSetter{Data: data, show: make([]*countryValue, 0), spacer: &layout.Spacer{}}
}

func (c *ContentSetter) sort(i int, j int) bool {
	return c.show[i].Overall >= c.show[j].Overall
}

func (c *ContentSetter) Resort() {
	sort.Slice(c.show, c.sort)
	c.updateWidgetContent()
}

func (c *ContentSetter) GetWidget() *container.Scroll {
	if c.widget == nil {
		c.widget = container.NewGridWithColumns(columns)
	}
	if c.scroll == nil {
		c.scroll = container.NewScroll(c.widget)
		c.scroll.OnScrolled = func(_ fyne.Position) {
			c.UpdateWidgetContent()
		}
	}
	return c.scroll
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
	if cap(c.parts) < len(c.show) {
		c.parts = make([]limits, len(c.show))
	}
	c.Resort()
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
	c.data = append(c.data,
		c.newHeader(headers[0]), c.newValue(values[0]))
	for hi := 1; hi < len(headers); hi++ {
		c.data = append(c.data,
			c.spacer, c.spacer,
			c.newHeader(headers[hi]), c.newValue(values[hi]))
	}
}

func (c *ContentSetter) updateWidgetContent() {
	c.GetWidget()
	if sz := c.requiredSize(); cap(c.data) < sz {
		c.data = make([]fyne.CanvasObject, 0, sz)
	}
	c.data = c.data[:0]
	c.parts = c.parts[:0]
	defer c.widget.Refresh()
	for i, country := range c.show {
		name := country.Name
		s := len(c.data)
		c.data = append(c.data, c.newTitle(i+1, name))

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
		c.data = append(c.data, cat)
		c.makeCategory(c.Data[category].Headers, values)
		for _, category = range country.Categories[hi+1:] {
			values = c.getValues(category, name)
			if len(values) == 0 {
				continue
			}
			c.data = append(c.data, c.spacer, c.newCategory(category))
			c.makeCategory(c.Data[category].Headers, values)
		}
		end := len(c.data)
		if end-s != 0 {
			c.parts = append(c.parts, limits{s, len(c.data)})
		}
	}
}
func (c *ContentSetter) getPartHeight(inx int) (h float32) {
	part := c.parts[inx]
	for i, item := range c.data[part.start:part.end] {
		if i%columns == 0 {
			h += item.Size().Height
		}
	}
	return
}
func (c *ContentSetter) UpdateWidgetContent() {
	pgSz := settings.Settings.PgSz
	if pgSz <= 0 {
		return
	}
	if len(c.widget.Objects) != 0 {
		if div := c.scroll.Content.Size().Height - c.scroll.Offset.Y; div <= c.scroll.Size().Height && c.curs+pgSz < len(c.parts) {
			c.curs++
			c.scroll.Offset.Y -= c.getPartHeight(c.curs + pgSz - 1)
		} else if c.curs > 0 && div >= c.scroll.Content.Size().Height {
			c.curs--
			c.scroll.Offset.Y += c.getPartHeight(c.curs)
		} else {
			return
		}
		c.widget.Objects = c.widget.Objects[:0]
	}
	e := c.curs + pgSz
	if e > len(c.parts) {
		e = len(c.parts)
	}
	for _, part := range c.parts[c.curs:e] {
		c.widget.Objects = append(c.widget.Objects, c.data[part.start:part.end]...)
	}
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
		c.Resort()
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
