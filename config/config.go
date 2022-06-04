package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadFile() *[]Account {
	f, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Print(err)
	}

	account := new([]Account)
	err = json.Unmarshal(f, &account)
	if err != nil {
		fmt.Print(err)
	}
	return account
}
