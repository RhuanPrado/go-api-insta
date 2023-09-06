package usermodule

import (
	userdto "go-api-insta/application/user/dto"
	"go-api-insta/libs/logger"
	"go-api-insta/models/api"

	"github.com/gofiber/fiber/v2"
)

func newRoutes(c Controller, app *fiber.App) {
	app.Post("/user", createUser(c))
	app.Put("/user", updatedUserName(c))
}

func createUser(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		user := &userdto.UserDto{}

		err := c.BodyParser(user)
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		err = user.Validate()

		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		response, statusCode, err := controller.CreateUser(user)

		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		return c.Status(statusCode).JSON(response)
	}
}
func updatedUserName(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//claims := jwt.GetClaims(c)
		id := ""
		user := &userdto.UserUpdateDto{}

		err := c.BodyParser(user)
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		err = user.Validate()

		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		response, statusCode, err := controller.UpdateUsername(id, user)

		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		return c.Status(statusCode).JSON(response)
	}
}
