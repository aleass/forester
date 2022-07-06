package config

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
