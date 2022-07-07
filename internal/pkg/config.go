package pkg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func InitConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	Conf := new(Config)
	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		return nil, err
	}
	return Conf, nil
}

type Config struct {
	Redis    Redis    `yaml:"redis"`
	Etcd     Etcd     `yaml:"etcd"`
	Http     Http     `yaml:"http"`
	ApiGrpc  Grpc     `yaml:"api_grpc"`
	TaskGrpc TaskGrpc `yaml:"task_grpc"`
}
type Redis struct {
	Addr string `yaml:"addr"`
}
type Etcd struct {
	Addr    string `yaml:"addr"`
	TaskPre string `yaml:"task_pre"`
}
type Http struct {
	Addr string `yaml:"addr"`
}

type Grpc struct {
	Addr string `yaml:"addr"`
}

type TaskGrpc struct {
	Port string `yaml:"port"`
}
