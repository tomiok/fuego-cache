package cache

import "fmt"

const (
	defaultPort         = 9919
	defaultMode         = "http"
	defaultFileLocation = "/home/fuego"
)

type FuegoConfig struct {
	//define if the cache should save the data in disk
	DiskPersistence bool
	// where save the data in the file system
	FileLocation string
	//http port
	WebPort string
	//mode - how to run the cache service, web, tco or CLI
	Mode string
}

func MakeConfig(persistence bool, fileLocation, mode string, port int) FuegoConfig {
	var (
		_fileLocation string
		_port         string
		_mode         string
	)

	if fileLocation == "" {
		_fileLocation = defaultFileLocation
	}

	if port == 0 {
		_port = fmt.Sprintf(":%d", defaultPort)
	}

	if mode == "" {
		_mode = defaultMode
	}

	return FuegoConfig{
		DiskPersistence: persistence,
		FileLocation:    _fileLocation,
		WebPort:         _port,
		Mode:            _mode,
	}
}
