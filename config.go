package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DatabaseLocation string `yaml:"database"`
	RootDir string `yaml:"root dir"`
	FoldersToIgnore []string `yaml:"folders to ignore"`
}

func OpenConfiguration(filename string) Config {
	var c Config
	config_raw, err := ioutil.ReadFile(filename)
	checkError(err)
	err = yaml.UnmarshalStrict(config_raw, &c)
	checkError(err)
	return c
}
