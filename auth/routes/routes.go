package routes

import (
	"auth-service/entity"
	"auth-service/services"

	"github.com/gofiber/fiber/v2"
)

func Auth(router fiber.Router) {
	router.Post("/token", generate())
	router.Get("/verify", verify())
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
		return nil
		// service := &services.Authenticator{}
		// fmt.Println(ctx.Request.Body())
		// service.GenerateToken(string(ctx.Request.Body()))
		// fmt.Fprintf(ctx, "Hello, world! Requested path is %q", ctx.Path())
	}
}
