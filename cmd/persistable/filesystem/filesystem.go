package filesystem

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type fileSystemStorage struct {
	rootPath string
}

func NewfileSystemStorage(rootPath string) *fileSystemStorage {
	log.Printf("New FileSystem storage targeting root path %s\n", rootPath)
	return &fileSystemStorage{rootPath}
}

func (f *fileSystemStorage) Save(b []byte, args ...any) error {
	if len(args) != 1 {
		return fmt.Errorf("wrong number of arguments received, expecting 1: <FILEPATH>")
	}

	filePath := args[0].(string)

	dirPath, fileName := extractDirPathFromFilePath(filePath)

	fmt.Println(dirPath, fileName)

	dirPath = filepath.Join(f.rootPath, dirPath)
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		return fmt.Errorf("can't create directories tree %s : %v", dirPath, err)
	}

	filePath = filepath.Join(dirPath, fileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("can't open the file %s : %v", filePath, err)
	}

	_, err = file.Write(b)
	if err != nil {
		return fmt.Errorf("can't write the file %s : %v", filePath, err)
	}

	log.Printf("File saved to %s\n", filePath)
	return nil
}

func (f *fileSystemStorage) Load(args ...any) ([]byte, error) {
	var data []byte
	if len(args) != 1 {
		return data, fmt.Errorf("wrong number of arguments received, expecting 1: <FILEPATH>")
	}

	filePath := args[0].(string)
	filePath = filepath.Join(f.rootPath, filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return data, fmt.Errorf("couldn't open file %s: %s", filePath, err)
	}

	data, err = io.ReadAll(file)
	if err != nil {
		return data, fmt.Errorf("couldn't read file %s: %s", filePath, err)
	}

	log.Printf("Loaded %s from filesystem\n", args)
	return data, nil
}

func extractDirPathFromFilePath(fp string) (string, string) {
	fpComponents := strings.Split(fp, "/")

	return strings.Join(fpComponents[:len(fpComponents)-1], "/"), fpComponents[len(fpComponents)-1]
}
