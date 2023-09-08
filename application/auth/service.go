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

	userL, err := a.userRepository.GetUserByUsername(data.Username)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error()}, fiber.StatusUnauthorized, err
	}

	match, err := argon2id.ComparePasswordAndHash(data.Password, userL.Password)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error()}, fiber.StatusUnauthorized, err
	}

	if match {
		dic := map[string]interface{}{
			"id": userL.ID.Hex(),
		}

		token, err := jwt.EncodeJwt(dic)
		if err != nil {
			return &api.Response{Error: true, ErrorMessage: err.Error()}, fiber.StatusInternalServerError, err
		}

		payload := struct {
			User struct {
				Id       string   `json:"id"`
				Username string   `json:"username"`
				Friends  []string `json:"friends"`
			} `json:"user"`
			Token string `json:"token"`
		}{}

		payload.User.Username = userL.Username
		payload.User.Id = userL.ID.Hex()
		payload.User.Friends = userL.Friends
		payload.Token = token

		return &api.Response{Error: false, Payload: payload}, fiber.StatusOK, err
	}

	return nil, fiber.StatusUnauthorized, nil
}
