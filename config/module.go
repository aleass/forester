package config

type Config struct {
	Redis Redis `yaml:"redis"`
	Etcd  Etcds `yaml:"etcd"`
}
type Redis struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
type Etcds struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}
