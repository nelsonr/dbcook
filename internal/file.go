package lib

import (
	"os"
	"path"
)

const (
	dirPerm  = 0755
	filePerm = 0644
)

func CreateFile(outputPath string, filename string, content string) error {
	err := os.MkdirAll(outputPath, dirPerm)
	if err != nil {
		return err
	}

	filePath := path.Join(outputPath, filename)
	err = os.WriteFile(filePath, []byte(content), filePerm)
	if err != nil {
		return err
	}

	return nil
}
