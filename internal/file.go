package lib

import (
	"os"
	"path/filepath"
	"strings"
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

	filePath := filepath.Join(outputPath, filename)
	err = os.WriteFile(filePath, []byte(content), filePerm)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func ListMigrationFiles(path string) ([]string, error) {
	filenames := []string{}
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".sql") {
			filenames = append(filenames, name)
		}
	}

	return filenames, nil
}
