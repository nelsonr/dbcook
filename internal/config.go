package lib

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const configFileName string = "dbcook.toml"

type Config struct {
	OutputPath string `toml:"output_path"`
}

func ReadConfigFile(path string) (Config, error) {
	var conf Config

	contents, err := os.ReadFile(filepath.Join(path, configFileName))
	if err != nil {
		return conf, err
	}

	conf, err = DecodeConfig(string(contents))
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func DecodeConfig(data string) (Config, error) {
	var conf Config

	_, err := toml.Decode(data, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
