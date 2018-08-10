package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"ioutil"
)

type Config struct {
	AccessLogFilePath string `yaml:"access_log_file_path"`
	ErrorLogFilePath  string `yaml:"error_log_file_path"`
	LogLevel          string `yaml:"log_level"`
	ServerID          string `yaml:"server_id"`
	BindHost          string `yaml:"bind_host"`
}

func main() {
	cfgPath := "./cfg.yaml"
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return
	}
}
