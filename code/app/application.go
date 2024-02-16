package app

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/masoud-mohajeri/kea-backend/config"
	"github.com/masoud-mohajeri/kea-backend/infra"
)

var app *fiber.App

func init() {

	var env string

	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	config.ParseEnv(env)
}

func StartApplication() {

	app = fiber.New(fiber.Config{
		UnescapePath: true,
		AppName:      "kea",
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			if e != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(e)
			}

			return nil
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // TODO: change it for prod
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, HEAD",
		AllowHeaders:     "Origin, Content-Length, Content-Type, Authorization, Accept",
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour), // for preflight reqs
	}))

	// compresses response
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// TODO: add sentry and logger
	infra.ConnectDB()
	infra.ConnectRedis()

	url_mapping()
	app.Listen(fmt.Sprintf(":%s", config.CommonConfig.Port))
}
