package config

import (
	"zapping.com/my_project/pkg/common"
)

type Config struct {
	myvar string
}

func New() *Config {
	return &Config{
		myvar: common.GetEnv("MYVAR", "default_value"),
	}
}
