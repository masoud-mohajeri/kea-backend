package routes

import (
	"github.com/gofiber/fiber/v2"
)

func NewPing(prefix string, r *fiber.App) {
	pinRouter := r.Group(prefix)

	pinRouter.Get("/", func(ctx *fiber.Ctx) error {
		ctx.JSON("PONG")
		return nil
	})
}
