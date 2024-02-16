package config

import (
	"errors"
	"os"
)

type common struct {
	Port string
}

var (
	CommonConfig = &common{}
)

func commonConfigModuleInit() {
	CommonConfig.Port = os.Getenv("PORT")

	err := CommonConfig.validation()

	if err != nil {
		panic(err)
	}
}

func (cfg *common) validation() error {
	if cfg.Port == "" {
		return errors.New("PORT environment required")
	}

	return nil
}
