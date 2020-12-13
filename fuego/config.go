package cache

import (
	"encoding/json"
	"os"
)

type FuegoConfig struct {
	//define if the cache should save the data in disk
	DiskPersistence bool `json:"disk_persistence"`
	// where save the data in the file system
	FileLocation string `json:"file_location"`
	//http port
	WebPort string `json:"web_port"`
	//mode - how to run the cache service, web, tcp or CLI
	Mode string `json:"mode"`
}

func ParseConfiguration() FuegoConfig {
	f, err := os.Open("config.json")

	defaultConfig := FuegoConfig{
		DiskPersistence: false,
		FileLocation:    "",
		WebPort:         "9919",
		Mode:            "http",
	}

	if err != nil {
		return defaultConfig
	}
	var fConfig FuegoConfig
	err = json.NewDecoder(f).Decode(&fConfig)

	if err != nil {
		return defaultConfig
	}

	return fConfig
}
