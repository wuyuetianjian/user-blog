package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Bootstrap struct {
	Server Server `yaml:"server"`
	MySQL  MySQL  `yaml:"mysql"`
}

type Server struct {
	HTTP HTTP `yaml:"http"`
}

type HTTP struct {
	Addr string `yaml:"addr"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

func Load(path string) (*Bootstrap, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Bootstrap
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
