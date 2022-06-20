package config

type Mysql struct {
	//公用配置
	Charset      string `yaml:"charset"`
	MaxIdleCoons string `yaml:"max_idle_coons"`
	MaxOpenCoons string `yaml:"max_open_coons"`
	//linkIt链接配置
	LinkItDb string `yaml:"link_it_db"`
	//系统库information_schema连接配置
	SchemaDb string `yaml:"schema_db"`
}
