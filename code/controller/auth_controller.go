package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/masoud-mohajeri/kea-backend/constants"
	"github.com/masoud-mohajeri/kea-backend/dto"
	"github.com/masoud-mohajeri/kea-backend/service"
)

type AuthController interface {
	OtpRequest(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	PasswordLogin(ctx *fiber.Ctx) error
	OtpLogin(ctx *fiber.Ctx) error
	ChangeMobile(ctx *fiber.Ctx) error
	// refresh token
}

type authController struct {
	otpService   service.OtpService
	userService  service.UserService
	tokenService service.TokenService
}

func NewAuthController(otpService service.OtpService, userService service.UserService, tokenService service.TokenService) AuthController {
	return &authController{
		otpService:   otpService,
		userService:  userService,
		tokenService: tokenService,
	}
}

func (ac *authController) OtpRequest(ctx *fiber.Ctx) error {
	req := new(dto.RequestOtpDto)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	vErr := req.Validate()
	if vErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(vErr.Error())
	}

	otp, osErr := ac.otpService.Request(req.Mobile)
	if osErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(osErr.Error())
	}

	user, urErr := ac.userService.GetUserByMobile(otp.Mobile)
	if urErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(urErr.Error())
	}
	isNewUser := false

	if user == nil {
		isNewUser = true
	}

	return ctx.Status(http.StatusOK).JSON(dto.OTPDetails{ExpiredAt: otp.ExpireAt, IsNew: isNewUser})
}

func (ac *authController) Register(ctx *fiber.Ctx) error {
	mobile := ctx.Params("mobile")
	if mobile == "" {
		return ctx.Status(http.StatusBadRequest).JSON("no mobile number")
	}

	body := new(dto.OtpValidate)
	if parsErr := ctx.BodyParser(body); parsErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON("body parsing error")
	}

	otpValidationErr := ac.otpService.Validate(body.Code, mobile)
	if otpValidationErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(otpValidationErr.Error())
	}

	saveErr := ac.userService.SaveUser(mobile, &body.UserInfo)
	if saveErr != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(saveErr.Error())
	}

	token, errT := ac.tokenService.CreateToken(mobile, constants.USER)
	if errT != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(errT.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(token)
}

func (ac *authController) PasswordLogin(ctx *fiber.Ctx) error {
	body := new(dto.PasswordLoginDto)

	parsErr := ctx.BodyParser(body)
	if parsErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON("body parsing error")
	}

	user, userErr := ac.userService.PasswordLogin(*body)

	if userErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(userErr.Error())
	}

	if user == nil {
		return ctx.Status(http.StatusNotFound).JSON("user not found")
	}

	role, roleErr := constants.GetRole(user.Role)
	if roleErr != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(roleErr.Error())
	}

	token, errT := ac.tokenService.CreateToken(body.Mobile, role)
	if errT != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(errT.Error())
	}

	return ctx.Status(http.StatusOK).JSON(token)
}

func (ac *authController) OtpLogin(ctx *fiber.Ctx) error {
	body := new(dto.OtpConfirmationDto)

	parsErr := ctx.BodyParser(body)
	if parsErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON("body parsing error")
	}

	err := ac.otpService.Validate(body.Otp, body.Mobile)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user, err := ac.userService.GetUserByMobile(body.Mobile)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	role, roleErr := constants.GetRole(user.Role)
	if roleErr != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(roleErr.Error())
	}

	token, errT := ac.tokenService.CreateToken(body.Mobile, role)
	if errT != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(errT.Error())
	}

	return ctx.Status(http.StatusOK).JSON(token)
}

func (ac *authController) ChangeMobile(ctx *fiber.Ctx) error {
	body := new(dto.ChangeMobileDto)

	parsErr := ctx.BodyParser(body)
	if parsErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON("body parsing error")
	}

	err := ac.otpService.Validate(body.Otp, body.Mobile)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user, err := ac.userService.UpdateMobile(body.Mobile, body.NewMobile)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	role, roleErr := constants.GetRole(user.Role)
	if roleErr != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(roleErr.Error())
	}

	token, errT := ac.tokenService.CreateToken(body.Mobile, role)
	if errT != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(errT.Error())
	}

	return ctx.Status(http.StatusOK).JSON(token)
}
