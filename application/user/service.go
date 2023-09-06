package usermodule

import (
	userdto "go-api-insta/application/user/dto"
	"go-api-insta/libs/logger"
	"go-api-insta/models/api"
	"go-api-insta/models/user"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(data *userdto.UserDto) (*api.Response, int, error)
	UpdateUsername(id string, data *userdto.UserUpdateDto) (*api.Response, int, error)
}

type userService struct {
	repository user.UserRepository
}

func newService() UserService {
	return &userService{}
}

func (u *userService) CreateUser(data *userdto.UserDto) (*api.Response, int, error) {

	newUser := u.userDtoToUser(data)

	hash, err := argon2id.CreateHash(data.Password, argon2id.DefaultParams)
	if err != nil {
		logger.Production.Info(err.Error())
		return &api.Response{Error: true, ErrorMessage: err.Error(), Payload: "error insert user"}, fiber.StatusInternalServerError, err
	}

	newUser.Password = hash

	_, err = u.repository.CreateUser(newUser)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error(), Payload: "error insert user"}, fiber.StatusInternalServerError, err
	}
	return &api.Response{Error: false, Payload: "success insert user"}, fiber.StatusOK, nil
}

func (u *userService) UpdateUsername(id string, data *userdto.UserUpdateDto) (*api.Response, int, error) {

	_, err := u.repository.UpdatedUserName(id, data.Username)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error(), Payload: "error insert user"}, fiber.StatusInternalServerError, err
	}
	return &api.Response{Error: false, Payload: "success insert user"}, fiber.StatusOK, nil
}

func (*userService) userDtoToUser(userDto *userdto.UserDto) *user.User {
	return &user.User{Username: userDto.Username, Password: userDto.Password}
}
