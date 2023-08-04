package guiNumbeo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/nikdissv-forever/numbeogo/gui/settings"
	"github.com/nikdissv-forever/numbeogo/gui/tables"
	"github.com/nikdissv-forever/numbeogo/recorder"
)

func Run() {
	recorder.SetRecord(true)
	win := app.New().NewWindow("Numbeo GUI")
	tabs := container.NewAppTabs()
	win.SetContent(tabs)
	settingsModal := settings.GetSettingsPopup(win.Canvas())
	win.SetMainMenu(
		fyne.NewMainMenu(fyne.NewMenu("File",
			fyne.NewMenuItem(
				"Settings", settingsModal.Show))))
	go processor(tabs)
	win.ShowAndRun()
}

func processor(tabs *container.AppTabs) {
	tabsTables := tables.TabsTables{Tabs: tabs}
	handler := func() {
		tabsTables.NewTablesByCategories()
	}
	settings.Signal.AddHandler(&handler)
	settings.Signal.Bell()
}
