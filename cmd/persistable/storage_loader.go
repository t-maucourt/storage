package persistable

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const (
	FILESYSTEM = "filesystem"
	S3         = "s3"
	SSH        = "ssh"
)

type Configuration struct {
	Storage storage `json:"storage"`
}

type storageSettings map[string]any

type storage struct {
	Type     string          `json:"type"`
	Settings storageSettings `json:"settings"`
}

func GetStorageFromConfiguration(configFilePath string) Persistable {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("can't open configuration file %s - %s", configFilePath, err)
	}

	defer configFile.Close()

	fileContent, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatalf("can't read the configuration file data - %s", err)
	}

	var config Configuration
	if err = json.Unmarshal(fileContent, &config); err != nil {
		log.Fatalf("can't unmarshal data from configuration file")
	}

	storageType := config.Storage.Type
	if storageType == "" {
		log.Fatalf("no storage type found in the configuration file")
	}

	var storageSystemBuilder func(storageSettings) Persistable
	switch storageType {
	case FILESYSTEM:
		storageSystemBuilder = NewfileSystemStorage
	default:
		log.Fatalf("storage system %s not found", config.Storage.Type)
	}

	return storageSystemBuilder(config.Storage.Settings)
}
