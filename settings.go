package tvtime

import (
	"encoding/json"
	"errors"

	"github.com/shibukawa/configdir"
)

var settingsDir = configdir.New("", "tvtime")

// Settings stores the access token
type Settings struct {
	AccessToken string
}

// HaveSettings returns true if we have some settings
func HaveSettings() bool {
	config := settingsDir.QueryFolderContainsFile("settings.json")
	return (config != nil)
}

// WriteSettings writes the settings
func WriteSettings(settings Settings) error {
	settingsData, err := json.Marshal(&settings)
	if err != nil {
		return err
	}

	settingsFolder := settingsDir.QueryFolders(configdir.Global)
	err = settingsFolder[0].WriteFile("settings.json", settingsData)
	return err
}

// GetSettings returns the settings
func GetSettings() (*Settings, error) {
	config := settingsDir.QueryFolderContainsFile("settings.json")

	if config == nil {
		return nil, errors.New("no settings found")
	}
	var settings Settings
	data, err := config.ReadFile("settings.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

// GetAccessToken returns the access token
func GetAccessToken() (string, error) {
	settings, err := GetSettings()
	if err != nil {
		return "", err
	}
	return settings.AccessToken, nil
}
