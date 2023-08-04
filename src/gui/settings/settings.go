package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/nikdissv-forever/numbeogo/internal/mutex"
	"github.com/nikdissv-forever/numbeogo/recorder"
	"sort"
)

const NoSelect = "Unselected"

type Storage struct {
	Title, Region string
}

var (
	Settings = Storage{}
)

func getSelectionListFromMap(src map[string]string) []string {
	mutex.Locker.Lock()
	defer mutex.Locker.Unlock()
	selection := make([]string, 0, len(src))
	for showValue := range src {
		selection = append(selection, showValue)
	}
	sort.Strings(selection)
	selection = append(selection, NoSelect)
	return selection
}

func getTitles() []string {
	return getSelectionListFromMap(recorder.Recorded.Titles)
}

func getRegions() []string {
	return getSelectionListFromMap(recorder.Recorded.Regions)
}

func getMapSelector(title string, bind *string, src map[string]string) (*widget.Label, *widget.Select) {
	selector := widget.NewSelect(getSelectionListFromMap(src),
		func(selected string) {
			if selected == NoSelect {
				*bind = ""
				Signal.Bell()
				return
			}
			newVal, ok := src[selected]
			if !ok || newVal == *bind {
				return
			}
			*bind = newVal
			Signal.Bell()
		},
	)
	selector.PlaceHolder = title
	selector.SetSelected(NoSelect)
	return widget.NewLabel(title), selector
}

func GetSettingsPopup(canvas fyne.Canvas) *widget.PopUp {
	titleLabel, titleSelector := getMapSelector("Year", &Settings.Title, recorder.Recorded.Titles)
	regionLabel, regionSelector := getMapSelector("Region", &Settings.Region, recorder.Recorded.Regions)
	h := func() {
		titleSelector.Options = getTitles()
		regionSelector.Options = getRegions()
	}
	recorder.Signal.AddHandler(&h)
	return widget.NewPopUp(container.NewVBox(container.NewPadded(container.NewGridWithColumns(2,
		titleLabel, regionLabel,
		titleSelector, regionSelector,
	))), canvas)
}
