package models

import (
	"os"
	"sync"

	"encoding/json"

	"github.com/adrg/xdg"
)

type Preferences struct {
	Theme         string
	Resolve       bool
	MaxHops       int
	Timeout       int
	ProbeCount    int
	JitterSamples int
}

var AppPreferencesMutex = sync.Mutex{}

func GetPreference(prefs *Preferences, name string) interface{} {
	AppPreferencesMutex.Lock()
	defer AppPreferencesMutex.Unlock()

	switch name {
	case "theme":
		return prefs.Theme
	case "resolve":
		return prefs.Resolve
	case "maxhops":
		return prefs.MaxHops
	case "timeout":
		return prefs.Timeout
	case "probecount":
		return prefs.ProbeCount
	case "jittersamples":
		return prefs.JitterSamples
	}
	return nil
}

func SetPreference(prefs *Preferences, name string, value interface{}) {
	AppPreferencesMutex.Lock()
	defer AppPreferencesMutex.Unlock()

	switch name {
	case "theme":
		prefs.Theme = value.(string)
	case "resolve":
		prefs.Resolve = value.(bool)
	case "maxhops":
		prefs.MaxHops = value.(int)
	case "timeout":
		prefs.Timeout = value.(int)
	case "probecount":
		prefs.ProbeCount = value.(int)
	case "jittersamples":
		prefs.JitterSamples = value.(int)
	}
}

func LoadPreferences(prefs *Preferences) error {
	AppPreferencesMutex.Lock()

	configPath, err := getXDGPath()
	if err != nil {
		AppPreferencesMutex.Unlock()
		return err
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// First time... default values
		prefs.Theme = "light"
		prefs.Resolve = false
		prefs.MaxHops = 30
		prefs.Timeout = 3
		prefs.ProbeCount = 3
		prefs.JitterSamples = 4
		AppPreferencesMutex.Unlock()
		SavePreferences(prefs)
		return nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		AppPreferencesMutex.Unlock()
		return err
	}
	if err := json.Unmarshal(data, prefs); err != nil {
		AppPreferencesMutex.Unlock()
		return err
	}
	AppPreferencesMutex.Unlock()
	return nil
}

func SavePreferences(prefs *Preferences) error {
	AppPreferencesMutex.Lock()
	defer AppPreferencesMutex.Unlock()

	configPath, err := getXDGPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(prefs, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return err
	}
	return nil
}

func getXDGPath() (string, error) {
	return xdg.ConfigFile("mtron/config.json")
}
