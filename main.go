package main

import (
	"fmt"
	"gofun/cmd/persistable"
	"log"
)

func main() {
	dataToSave := []byte("Hello World")
	filePath := "folder_1/folder_2/some-file.json"

	s := persistable.GetStorageFromConfiguration("config.json")

	if err := s.Save(dataToSave, filePath); err != nil {
		log.Printf("Failed to save data: %s\n", err)
	}

	data, err := s.Load(filePath)
	if err != nil {
		log.Printf("Failed to load data: %s\n", err)
	}

	fmt.Println(string(data))
}
