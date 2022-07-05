package config

import (
	"Forester/config"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

func InitConfig(path string) (*config.Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	Conf := new(config.Config)
	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		return nil, err
	}
	return Conf, nil
}
