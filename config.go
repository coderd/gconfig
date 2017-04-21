package gconfig

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigFile struct {
	file string
	data map[string]interface{}
}

func LoadJsonFile(file string) (*ConfigFile, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	configFile := &ConfigFile{
		file: file,
		data: data,
	}

	return configFile, nil
}
