package config

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Auth string `yaml:"auth"`
	Db   string `yaml:"db"`
}
