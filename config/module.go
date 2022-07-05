package config

type Config struct {
	Redis Redis `yaml:"redis"`
	Etcd  Etcd  `yaml:"etcd"`
	Http  Http  `yaml:"http"`
}
type Redis struct {
	Addr string `yaml:"addr"`
}
type Etcd struct {
	Addr string `yaml:"addr"`
}
type Http struct {
	Addr string `yaml:"addr"`
}
