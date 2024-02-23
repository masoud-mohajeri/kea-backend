package config

import (
	"errors"
	"os"
	"time"
)

type common struct {
	Port                 string
	TokenSecret          []byte
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

var (
	CommonConfig = &common{}
)

func commonConfigModuleInit() {
	CommonConfig.Port = os.Getenv("PORT")
	CommonConfig.TokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
	CommonConfig.AccessTokenDuration, _ = time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	CommonConfig.RefreshTokenDuration, _ = time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))

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
