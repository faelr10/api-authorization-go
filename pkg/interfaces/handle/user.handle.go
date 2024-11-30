package user_handler_pkg

import (
	"github.com/gofiber/fiber/v2"
)

type IUserHandlers interface {
	NewUser(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
}
