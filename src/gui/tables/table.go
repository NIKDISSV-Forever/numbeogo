package tables

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nikdissv-forever/numbeogo/gui/settings"
	"github.com/nikdissv-forever/numbeogo/gui/settings/save"
	"github.com/nikdissv-forever/numbeogo/parser"
	"golang.org/x/image/colornames"
	"strconv"
	"strings"
)

type row struct {
	name   string
	values []float32
}

type OrderedTable struct {
	headers []string
	rows    row
}
type TabsTables struct {
	Tabs *container.AppTabs
}

type TableView struct {
	*TabsTables
	Tab *container.TabItem

	updates chan bool
	table   numbeo.Table
	root    *fyne.Container
}

func (s *TableView) Refresh() {
	s.Tabs.Refresh()
	s.Tab.Content.Refresh()
}

func (s *TableView) TableController(url string) {
	s.Tab.Icon = theme.ViewRefreshIcon()
	s.table = numbeo.GetPageTable(numbeo.TableParams{
		URL:    url,
		Title:  settings.Settings.Title,
		Region: settings.Settings.Region,
	})
	s.updates = make(chan bool, 1)
	s.updates <- true
	halt := func() {
		s.updates <- false
	}
	settings.Signal.AddHandler(&halt)
	for {
		select {
		case v := <-s.updates:
			if v {
				s.makeTable()
				s.Tab.Content = container.NewVScroll(container.NewHScroll(container.NewPadded(s.root)))
				s.Tab.Icon = theme.ListIcon()
				s.Refresh()
			} else {
				halt = nil
				return
			}
		}
	}
}

func (s *TableView) makeTable() {
	sortTable(s.table)
	columnsCount := len(s.table.Headers) + 1
	s.root = container.NewGridWithColumns(columnsCount)
	s.root.Objects = make([]fyne.CanvasObject, 0, columnsCount*(len(s.table.CountryValues)+1))
	s.pasteHeaders()
	s.pasteRows()
}

func (s *TableView) appendObject(object fyne.CanvasObject) {
	s.root.Objects = append(s.root.Objects, object)
}

func (s *TableView) pasteHeaders() {
	saveNew := false
	s.appendObject(container.NewCenter(canvas.NewText("Country", colornames.Brown)))
	for _, header := range s.table.Headers {
		isNegative := false
		if isNeg, ok := save.Settings.Sorts[header]; ok {
			isNegative = isNeg
		} else {
			save.Settings.Sorts[header] = false
			saveNew = true
		}
		var colorName widget.ButtonImportance
		if isNegative {
			colorName = widget.LowImportance
		} else {
			colorName = widget.HighImportance
		}
		btn := widget.NewButton(header, s.onTapped(header))
		btn.Importance = colorName
		s.appendObject(btn)
	}
	if saveNew {
		save.SettingsSave()
	}
}

func (s *TableView) pasteRows() {
	for i := range s.table.CountryValues {
		s.pasteRow(i)
	}
}

func (s *TableView) pasteRow(rowIndex int) {
	country := s.table.CountryValues[rowIndex]
	s.appendObject(container.NewHBox(
		canvas.NewText(strconv.Itoa(rowIndex+1), colornames.Blue),
		canvas.NewText(country.Name, colornames.Brown),
	))
	for _, value := range country.Values {
		s.appendObject(canvas.NewText(fmt.Sprintf("%g", value), colornames.Blueviolet))
	}
}

func (s *TableView) onTapped(header string) func() {
	return func() {
		save.Settings.Sorts[header] = !save.Settings.Sorts[header]
		s.updates <- true
		save.SettingsSave()
	}
}

func (s *TabsTables) NewTablesByCategories() {
	categories, err := numbeo.GetRankingByCountryCategories()
	if err != nil {
		return
	}
	items := make([]*container.TabItem, 0, len(categories))
	for _, cat := range categories {
		tab := container.NewTabItemWithIcon(cat.Name, theme.DownloadIcon(),
			container.NewCenter(
				canvas.NewText(strings.Join([]string{"Loading ", cat.Name, "..."}, ""),
					colornames.Green)))
		items = append(items, tab)
		s.Tabs.SetItems(items)
		table := TableView{TabsTables: s, Tab: tab}
		go table.TableController(cat.URL)
	}
	s.Tabs.Items = items
}
