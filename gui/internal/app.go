package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var RootApp = app.NewWithID("com.numbeo")
var MainWindow = RootApp.NewWindow("Numbeo GUI")
var DefaultSize = fyne.Size{Width: 640, Height: 360}

func init() {
	MainWindow.SetMaster()
}
