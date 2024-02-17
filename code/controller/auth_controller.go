package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/masoud-mohajeri/kea-backend/dto"
	"github.com/masoud-mohajeri/kea-backend/service"
)

type AuthController interface {
	OtpRequest(ctx *fiber.Ctx) error
	// Login(ctx *fiber.Ctx) error
	// PasswordLogin(ctx *fiber.Ctx) error
	// ChangePhoneNumber(ctx *fiber.Ctx) error
}

type authController struct {
	otpService  service.OtpService
	userService service.UserService
}

func NewAuthController(otpService service.OtpService, userService service.UserService) AuthController {
	return &authController{
		otpService:  otpService,
		userService: userService,
	}
}

func (ac *authController) OtpRequest(ctx *fiber.Ctx) error {
	req := new(dto.RequestOtpDto)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	vErr := req.Validate()
	if vErr != nil {
		return vErr
	}

	otp, osErr := ac.otpService.Request(req.Mobile)
	if osErr != nil {
		return osErr
	}

	user, urErr := ac.userService.GetUserByMobile(otp.Mobile)

	if urErr != nil {
		return urErr
	}
	isNewUser := false

	if user == nil {
		isNewUser = true
	}

	return ctx.Status(http.StatusOK).JSON(dto.OTPDetails{ExpiredAt: otp.ExpireAt, IsNew: isNewUser})
}
