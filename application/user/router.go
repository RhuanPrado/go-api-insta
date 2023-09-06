package usermodule

import (
	userdto "go-api-insta/application/user/dto"
	"go-api-insta/libs/jwt"
	"go-api-insta/libs/logger"
	"go-api-insta/models/api"

	"github.com/gofiber/fiber/v2"
)

func newRoutes(c Controller, app *fiber.App) {
	app.Post("/user", createUser(c))
	app.Put("/user", jwt.JwtProtected(), updatedUserName(c))
	app.Put("/user/friends", jwt.JwtProtected(), updateFriends(c))
	app.Get("/users", jwt.JwtProtected(), findUsers(c))
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
		id := jwt.DecodeJwtSingleKey(c, "id").(string)
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

func updateFriends(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		id := jwt.DecodeJwtSingleKey(c, "id").(string)
		user := &userdto.UserUpdateFriendsDto{}

		err := c.BodyParser(user)
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		response, statusCode, err := controller.UpdateUserFriends(id, user.Friends)
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

func findUsers(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := jwt.DecodeJwtSingleKey(c, "id").(string)

		response, statusCode, err := controller.FindUsers(id)

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
