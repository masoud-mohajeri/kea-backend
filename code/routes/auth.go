package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masoud-mohajeri/kea-backend/controller"
	"github.com/masoud-mohajeri/kea-backend/infra"
	"github.com/masoud-mohajeri/kea-backend/repository"
	"github.com/masoud-mohajeri/kea-backend/service"
)

func NewAuth(prefix string, r *fiber.App) {
	redisRepository := repository.NewRedisRepository(infra.RedisClient)
	otpRepository := repository.NewOtpRepository(redisRepository)
	smsService := service.NewSmsService()
	otpService := service.NewOtpService(smsService, otpRepository)
	userRepository := repository.NewUserRepository(infra.DB)
	authController := controller.NewAuthController(otpService, userRepository)

	routes := r.Group(prefix)

	routes.Post("otp-request", authController.OtpRequest)
	// routes.Post("register")
	// routes.Post("otp-login")
	// routes.Post("password-login")
	// routes.Post("change-phone")

}