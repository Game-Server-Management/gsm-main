package config

import (
	"encoding/json"
	"gsm/utilities"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ServerInfo struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

type Config struct {
	Servers    map[string]ServerInfo `json:"servers"`
	Games_path string                `json:"installation_path"`
}

var Name = "Game Server Management Utility"
var Configuration = Config{}
var ConfigPath = filepath.Join("/etc", "gsm", "config.json")

func Init() {
	Configuration.Servers = make(map[string]ServerInfo, 0)
	ReadConfig(ConfigPath)
}

func ReadConfig(path string) {
	if utilities.FileExists(path) {
		config, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(config, &Configuration)
		if err != nil {
			panic(err)
		}
	}
}

func SaveConfig(path string) {
	content, err := json.MarshalIndent(Configuration, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(path), 0644)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, content, 0644)
	if err != nil {
		panic(err)
	}
}
