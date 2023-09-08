package postmodule

import (
	postdto "go-api-insta/application/post/dto"
	"go-api-insta/libs/jwt"
	"go-api-insta/libs/logger"
	"go-api-insta/models/api"
	"io"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func newRoutes(c Controller, app *fiber.App) {

	app.Post("/post", jwt.JwtProtected(), createPost(c))
	app.Get("/post/all", jwt.JwtProtected(), findAllPostFriends(c))
	app.Get("/post/:userId", jwt.JwtProtected(), getPostsUser(c))
}

func createPost(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := jwt.DecodeJwtSingleKey(c, "id").(string)
		post := &postdto.PostDto{}

		file, err := c.FormFile("file")
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}
		f, _ := file.Open()
		defer f.Close()

		post.File, err = readFileBytes(f)
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		post.Description = c.FormValue("description")

		err = post.Validate()
		if err != nil {
			logger.Production.Info(err.Error())
			c.Status(fiber.StatusBadRequest)
			return c.JSON(api.Response{
				Error:        true,
				ErrorMessage: err.Error(),
			})
		}

		response, statusCode, err := controller.CreatePost(id, post)
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

func getPostsUser(controller Controller) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		postUser := &postdto.PostUserDto{}

		postUser.UserId = c.Params("userId")

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

func readFileBytes(file multipart.File) ([]byte, error) {
	var fileBytes []byte
	buffer := make([]byte, 1024)

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if bytesRead == 0 {
			break
		}
		fileBytes = append(fileBytes, buffer[:bytesRead]...)
	}

	return fileBytes, nil
}
