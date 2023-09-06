package authmodule

import (
	authdto "go-api-insta/application/auth/dto"
	"go-api-insta/libs/jwt"
	"go-api-insta/models/api"
	"go-api-insta/models/user"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Authorization(data *authdto.AuthDto) (*api.Response, int, error)
}

type authService struct {
	userRepository user.UserRepository
}

func newService() AuthService {
	return &authService{}
}

func (a *authService) Authorization(data *authdto.AuthDto) (*api.Response, int, error) {

	user, err := a.userRepository.GetUserByUsername(data.Username)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error()}, fiber.StatusBadRequest, err
	}

	match, err := argon2id.ComparePasswordAndHash(data.Password, user.Password)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error()}, fiber.StatusUnauthorized, err
	}
	println(user.ID.String())
	if match {
		dic := map[string]interface{}{
			"id": user.ID.String(),
		}

		token, err := jwt.EncodeJwt(dic)
		if err != nil {
			return &api.Response{Error: true, ErrorMessage: err.Error()}, fiber.StatusInternalServerError, err
		}

		return &api.Response{Error: false, Payload: token}, fiber.StatusOK, err
	}

	return nil, fiber.StatusUnauthorized, nil
}
