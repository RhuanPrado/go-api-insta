package postmodule

import (
	postdto "go-api-insta/application/post/dto"
	"go-api-insta/libs/jwt"
	"go-api-insta/libs/logger"
	"go-api-insta/models/api"

	"github.com/gofiber/fiber/v2"
)

func newRoutes(c Controller, app *fiber.App) {
	app.Post("/post", jwt.JwtProtected(), createPost(c))
	app.Get("/post", jwt.JwtProtected(), getPostsUser(c))
	app.Get("/post/all", jwt.JwtProtected(), findAllPostFriends(c))
}

func createPost(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := jwt.DecodeJwtSingleKey(c, "id").(string)
		post := &postdto.PostDto{}

		err := c.BodyParser(post)
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		controller.CreatePost(id, post)
		return nil
	}
}

func getPostsUser(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		postUser := &postdto.PostUserDto{}

		err := c.BodyParser(postUser)
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}
		response, statusCode, err := controller.FindAllPostUser(postUser)
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

func findAllPostFriends(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		id := jwt.DecodeJwtSingleKey(c, "id").(string)

		response, statusCode, err := controller.FindAllPostFriends(id)
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
