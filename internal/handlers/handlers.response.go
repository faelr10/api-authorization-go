package handlers

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewSuccessResponse(ctx *fiber.Ctx, statusCode int, message string, payload interface{}) error {
	return ctx.Status(statusCode).JSON(BaseResponse{
		Status:  true,
		Message: message,
		Payload: payload,
	})
}

func NewErrorReponse(ctx *fiber.Ctx, statusCode int, err string) error {
	return ctx.Status(statusCode).JSON(BaseResponse{
		Status:  false,
		Message: err,
	})
}
