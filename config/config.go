package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadConfigFile(filePath string) (Config, error) {
	var config Config

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("failed to read file: %w", err)
	}

	if err = json.Unmarshal(bytes, &config); err != nil {
		return config, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return config, nil
}
