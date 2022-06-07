package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadFile() Config {
	f, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Print(err)
	}

	var config Config
	err = json.Unmarshal(f, &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
