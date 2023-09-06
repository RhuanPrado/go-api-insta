package jwt

import (
	"go-api-insta/helpers/variable"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	jwtMiddleware "github.com/gofiber/jwt/v3"
)

func jwtError(c *fiber.Ctx, err error) error {

	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": true,
			"msg": err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"err": true,
		"msg": err.Error(),
	})
}

func JwtProtected() func(*fiber.Ctx) error {

	config := jwtMiddleware.Config{
		SigningKey:   variable.GetEnvVariable("JWT_KEY"),
		ContextKey:   "authorization",
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func GetClaims(c *fiber.Ctx) jwt.MapClaims {

	// Parses the JWT used to secure authorized access to private routes
	token := c.Locals("authorization").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Forwards the claims further to the function that is using them
	return claims
}
