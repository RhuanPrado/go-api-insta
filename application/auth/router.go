package authmodule

import (
	authdto "go-api-insta/application/auth/dto"
	"go-api-insta/libs/logger"
	"go-api-insta/models/api"

	"github.com/gofiber/fiber/v2"
)

func newRoutes(c Controller, app *fiber.App) {
	app.Post("/authorization", authorization(c))
}

func authorization(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		auth := &authdto.AuthDto{}

		err := c.BodyParser(auth)
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		err = auth.Validate()
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		response, statusCode, err := controller.Authorization(auth)
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(statusCode)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		return c.Status(statusCode).JSON(response)
	}
}
