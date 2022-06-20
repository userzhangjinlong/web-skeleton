package config

//接收制定yaml配置
type Frame struct {
	Debug string `yaml:"debug"`
	Port  string `yaml:"port"`
}

type Log struct {
	Path string `yaml:"path"`
}
