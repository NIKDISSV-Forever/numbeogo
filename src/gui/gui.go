package guiNumbeo

import (
	"fyne.io/fyne/v2"
	"github.com/nikdissv-forever/numbeogo/gui/internal"
	"github.com/nikdissv-forever/numbeogo/gui/settings"
	numbeo "github.com/nikdissv-forever/numbeogo/parser"
	"github.com/nikdissv-forever/numbeogo/recorder"
	"sync"
)

func Run() {
	recorder.SetRecord(true)
	paramsPopup := settings.GetSettingsPopup()
	setter := NewContentSetter(make(map[string]*numbeo.Table))
	internal.MainWindow.SetMainMenu(
		fyne.NewMainMenu(fyne.NewMenu("Settings",
			fyne.NewMenuItem("Parameters", paramsPopup.Show),
			settings.NewCountriesChoicer().GetChoiceCountriesMenuOption(func() { setter.Refresh() }))))
	internal.MainWindow.SetContent(setter.GetWidget())
	handler := func() {
		setter.Data = getData()
		setter.UpdateData()
		setter.Resort()
		setter.UpdateWidgetContent()
	}
	settings.Signal.AddHandler(&handler)
	go settings.Signal.Bell()
	internal.MainWindow.Resize(internal.DefaultSize)
	internal.MainWindow.ShowAndRun()
}

func getData() map[string]*numbeo.Table {
	categories, err := numbeo.GetRankingByCountryCategories()
	if err != nil {
		return make(map[string]*numbeo.Table)
	}
	data := make(map[string]*numbeo.Table, len(categories))
	wg := sync.WaitGroup{}
	wg.Add(len(categories))
	for _, cat := range categories {
		go func(cat *numbeo.Category) {
			defer wg.Done()
			table := numbeo.GetPageTable(numbeo.TableParams{
				URL:    cat.URL,
				Title:  settings.Settings.Title,
				Region: settings.Settings.Region,
			})
			data[cat.Name] = table
		}(cat)
	}
	wg.Wait()
	return data
}
