package configs

import (
	"log"
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func NewConfig() *koanf.Koanf {

	koanf := koanf.New(".")

	err := koanf.Load(file.Provider(checkConfigPath()), yaml.Parser())
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	return koanf

}

func checkConfigPath() string {

	location := []string{
		"configs/conf.yaml",
		"./../configs/conf.yaml",
		"./../../configs/conf.yaml",
		"./../../../configs/conf.yaml",
	}
	for _, location := range location {
		_, err := os.Stat(location)
		if err == nil {
			return location
		}

	}
	return ""

}
