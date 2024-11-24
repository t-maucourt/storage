package main

import (
	"fmt"
	"gofun/cmd/persistable"
)

func main() {
	dataToSave := []byte("Hello World")
	filePath := "folder_1/folder_2/some-file.json"

	s := persistable.GetStorageFromConfiguration("config.json")

	err := s.Save(dataToSave, filePath)
	fmt.Println(err)
	data, err := s.Load(filePath)
	fmt.Println(err)
	fmt.Println(string(data))
}
