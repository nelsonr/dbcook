package lib

import (
	"os"
	"path"
)

const (
	PERM_DIR  = 0755
	PERM_FILE = 0644
)

func CreateFile(dirPath string, filename string, content string) error {
	err := os.MkdirAll(dirPath, PERM_DIR)
	if err != nil {
		return err
	}

	filePath := path.Join(dirPath, filename)
	err = os.WriteFile(filePath, []byte(content), PERM_FILE)
	if err != nil {
		return err
	}

	return nil
}
