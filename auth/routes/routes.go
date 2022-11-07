package routes

import (
	"auth-service/entity"
	"auth-service/services"

	"github.com/gofiber/fiber/v2"
)

func Auth(router fiber.Router) {
	router.Post("/token", generate())
	router.Post("/verify", verify())
}

func generate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entity.UserAuth
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.JSON(&fiber.Map{
				"status": fiber.StatusBadRequest,
				"data":   "",
				"error":  err.Error(),
			})
		}

		service := services.New()

		jwtToken, err := service.GenerateToken(requestBody.Username, requestBody.Password)
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.JSON(&fiber.Map{
				"status": fiber.StatusBadRequest,
				"data":   "",
			})
		}
		return c.JSON(&fiber.Map{
			"status": fiber.StatusOK,
			"token":  jwtToken,
		})
	}
}

func verify() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entity.TokenEntity
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.JSON(&fiber.Map{
				"status": fiber.StatusBadRequest,
				"data":   "",
				"error":  err.Error(),
			})
		}

		service := services.New()

		payload, err := service.VerifyToken(requestBody)
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.JSON(&fiber.Map{
				"status": fiber.StatusBadRequest,
				"data":   "",
			})
		}
		return c.JSON(&fiber.Map{
			"status":  fiber.StatusOK,
			"payload": payload,
		})
	}
}
