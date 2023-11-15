package settings

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hashicorp/go-set"
	"github.com/nikdissv-forever/numbeogo/gui/internal"
	customWidgets "github.com/nikdissv-forever/numbeogo/gui/pkg/widget"
	"github.com/nikdissv-forever/numbeogo/gui/resources"
	"github.com/nikdissv-forever/numbeogo/gui/settings/preferences/settings"
	"github.com/nikdissv-forever/numbeogo/internal/mutex"
	"github.com/nikdissv-forever/numbeogo/recorder"
	"sort"
	"strconv"
)

const NoSelect = "Unselected"

type Storage struct {
	Countries     set.Set[string]
	Title, Region string
	PgSz          int
}

var Settings = Storage{PgSz: settings.GetPgSz()}

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

func getPgSizeEntry(title string) (*widget.Label, *customWidgets.NumericalEntry) {
	entry := customWidgets.NewNumericalEntry()
	entry.SetPlaceHolder(title)
	entry.OnSubmitted = func(s string) {
		i, err := strconv.Atoi(s)
		if err != nil {
			entry.SetText(strconv.Itoa(Settings.PgSz))
			return
		}
		Settings.PgSz = i
		settings.SetPgSz(i)
	}
	return widget.NewLabel(title), entry
}

func GetSettingsPopup() *widget.PopUp {
	titleLabel, titleSelector := getMapSelector("Year", &Settings.Title, recorder.Recorded.Titles)
	regionLabel, regionSelector := getMapSelector("Region", &Settings.Region, recorder.Recorded.Regions)
	pgSzLabel, pgSzSelector := getPgSizeEntry("Page Size")
	pgSzSelector.SetText(strconv.Itoa(Settings.PgSz))
	h := func() {
		titleSelector.Options = getTitles()
		regionSelector.Options = getRegions()
	}
	recorder.Signal.AddHandler(&h)
	return widget.NewPopUp(container.NewPadded(container.NewGridWithColumns(2,
		titleLabel, titleSelector,
		regionLabel, container.NewHBox(regionSelector, setMyRegion(regionSelector)),
		pgSzLabel, pgSzSelector,
	)), internal.MainWindow.Canvas())
}

func setMyRegion(selector *widget.Select) *widget.Button {
	btn := widget.NewButtonWithIcon("", resources.MapMarker, func() {
		if reg, err := getMyRegion(); err == nil {
			selector.SetSelected(reg)
		}
	})
	btn.Importance = widget.LowImportance
	return btn
}
