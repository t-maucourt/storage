package persistable

import (
	"encoding/json"
	"gofun/cmd/persistable/filesystem"
	"gofun/cmd/persistable/s3"
	"io"
	"log"
	"os"
	"strings"
)

const (
	FILESYSTEM = "filesystem"
	S3         = "s3"
)

type configuration struct {
	Storage map[string]string `json:"storage"`
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
	var config configuration

	if err = json.Unmarshal(fileContent, &config); err != nil {
		log.Fatalf("can't unmarshal data from configuration file")
	}

	storageType := config.Storage["type"]
	if storageType == "" {
		log.Fatalf("no storage type found in the configuration file")
	}

	storageType = strings.ToLower(storageType)

	var storageSystem Persistable
	switch storageType {
	case FILESYSTEM:
		storageSystem = filesystem.NewfileSystemStorage(config.Storage["root_path"])
	case S3:
		storageSystem = s3.NewS3Storage(config.Storage["bucket"])
	default:
		log.Fatalf("storage system %s not found", config.Storage["type"])
	}

	return storageSystem

}
