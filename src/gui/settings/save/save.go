package save

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"os"
)

const settingsPath = "settings.json"

func init() {
	if stat, err := os.Stat(settingsPath); err == nil && stat.Size() > 0 {
		if file, err := os.ReadFile(settingsPath); err == nil {
			newSetting := SettingsType{}
			if err = json.Unmarshal(file, &newSetting); err == nil {
				Settings = newSetting
			}
		}
	} else {
		SettingsSave()
	}
}

func SettingsSave() {
	marshaled, err := json.MarshalIndent(Settings, "", " ")
	if err != nil {
		fyne.LogError("SettingsSave", err)
		return
	}

	if err = os.WriteFile(settingsPath, marshaled, 0666); err != nil {
		return
	}
}
