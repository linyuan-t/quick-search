package config

type server struct {
	ID    int64  `yaml:"id"`
	Debug bool   `yaml:"debug"`
	Port  string `yaml:"port"`
}

// Config 配置写这里，传递到各个package
type Config struct {
	FileName string
	Server   server `yaml:"server"`
}
