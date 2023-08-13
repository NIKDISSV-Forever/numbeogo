package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hashicorp/go-set"
	"github.com/nikdissv-forever/numbeogo/gui/internal"
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences/countriesFilter"
	"github.com/nikdissv-forever/numbeogo/recorder"
	"math"
	"regexp"
	"sort"
)

type CountriesChoicer struct {
	filter func(string) bool
	root   *fyne.CanvasObject
	showed *set.Set[string]
}

func NewCountriesChoicer() *CountriesChoicer {
	return &CountriesChoicer{filter: regexp.MustCompile(".*").MatchString}
}

func (s *CountriesChoicer) GetChoiceCountriesMenuOption(onClosed func()) *fyne.MenuItem {
	item := fyne.NewMenuItem("Countries", func() {
		w := internal.RootApp.NewWindow("Counties")
		w.SetOnClosed(onClosed)
		w.Resize(internal.DefaultSize)
		w.CenterOnScreen()
		w.SetContent(s.getChoicerContent())
		w.Show()
	})
	if recorder.Recorded.Countries.Size() == 0 {
		item.Disabled = true
		go func() {
			for recorder.Recorded.Countries.Size() == 0 {
			}
			item.Disabled = false
		}()
	}
	return item
}

func (s *CountriesChoicer) getChoicerContent() *container.Scroll {
	box := container.NewVBox()
	box.Objects = make([]fyne.CanvasObject, 2)
	box.Objects[0] = s.getControlButtons()
	s.root = &box.Objects[1]
	s.getChoicerTable()
	return container.NewHScroll(container.NewVScroll(box))
}

func (s *CountriesChoicer) getControlButtons() *fyne.Container {
	selectAll := widget.NewCheck("Show", s.SetAll)
	selectAll.Checked = true
	search := widget.NewEntry()
	search.Wrapping = fyne.TextWrapOff
	search.PlaceHolder = "Search"
	search.OnChanged = s.ApplyFilter
	return container.NewHBox(selectAll, search)
}

func (s *CountriesChoicer) ApplyFilter(query string) {
	s.filter = internal.GetFilterFunc(query)
	s.getChoicerTable()
}

func (s *CountriesChoicer) SetAll(b bool) {
	s.showed.ForEach(func(name string) bool {
		countriesFilter.Set(name, b)
		return true
	})
	s.getChoicerTable()
}

func (s *CountriesChoicer) getChoicerTable() {
	s.showed = recorder.Recorded.Countries.Copy()
	s.showed.RemoveFunc(func(q string) bool { return !s.filter(q) })
	list := s.showed.Slice()
	sort.Strings(list)
	lines := make([]fyne.CanvasObject, 0, len(list))
	for _, country := range list {
		lines = append(lines, s.getLine(country))
	}
	*s.root = container.NewGridWithRows(int(math.Ceil(math.Sqrt(float64(len(lines))))), lines...)
}

func (s *CountriesChoicer) getLine(country string) *widget.Check {
	check := widget.NewCheck(country, func(b bool) {
		countriesFilter.Set(country, b)
	})
	check.Checked = countriesFilter.Get(country)
	return check
}
