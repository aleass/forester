package config

import (
	"Forester/config"
	"Forester/internal/manager"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

func InitConfig(server *manager.Server, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	Conf := new(config.Config)
	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		return err
	}
	server.Config = Conf
	return nil
}
