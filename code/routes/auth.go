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
	userService := service.NewUserService(userRepository)
	tokenService := service.NewTokenService()
	authController := controller.NewAuthController(otpService, userService, tokenService)

	routes := r.Group(prefix)

	routes.Post("otp-request", authController.OtpRequest)
	routes.Post("register/:mobile", authController.Register)
	routes.Post("password-login", authController.PasswordLogin)
	routes.Post("otp-login", authController.OtpLogin)
	routes.Post("refresh", authController.RefreshToken)
	// TODO: move it to user routes
	routes.Post("change-mobile", authController.ChangeMobile)

}
