package config

import (
	"github.com/joho/godotenv"
)

func ParseEnv(e string) {
	var fn string

	if e == "prod" {
		fn = "prod.env"
	} else {
		fn = "local.env"
	}

	err := godotenv.Load(fn)

	if err != nil {
		godotenv.Load()
	}

	commonConfigModuleInit()
	databaseConfigModuleInit()
	redisConfigModuleInit()
}
