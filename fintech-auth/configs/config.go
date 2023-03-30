package configs

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func NewConfig() *koanf.Koanf {

	koanf := koanf.New(".")

	err := koanf.Load(file.Provider("configs/conf.yaml"), yaml.Parser())
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	return koanf

}
