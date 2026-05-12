package lib

import (
	"os"
	"path"
)

func CreateFile(dirPath string, filename string, content string) error {
	err := createDir(dirPath)
	if err != nil {
		return err
	}

	filePath := path.Join(dirPath, filename)
	err = createFile(filePath, content)
	if err != nil {
		return err
	}

	return nil
}

func createDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

func createFile(filePath string, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}
