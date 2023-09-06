package main

import (
	"go-api-insta/helpers/variable"
	"go-api-insta/libs/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	loggerFiber "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	logger.InitializeLogger()

	app := fiber.New()

	app.Use(recover.New())

	app.Use(helmet.New())

	app.Use(cors.New(cors.ConfigDefault))

	app.Use(loggerFiber.New())

	app.Listen(":" + variable.GetEnvVariable("PORT"))

}
