package usermodule

import (
	userdto "go-api-insta/application/user/dto"
	"go-api-insta/models/api"
)

// Controller defines the methods to be exposed in Controller layer
type Controller interface {
	CreateUser(data *userdto.UserDto) (*api.Response, int, error)
	UpdateUsername(id string, data *userdto.UserUpdateDto) (*api.Response, int, error)
	FindUsers(userId string) (*api.Response, int, error)
	UpdateUserFriends(id string, friends []string) (*api.Response, int, error)
}

type controller struct {
	UserService
}

func newController(service UserService) Controller {
	return &controller{
		UserService: service,
	}
}

func (c *controller) CreateUser(data *userdto.UserDto) (*api.Response, int, error) {
	return c.UserService.CreateUser(data)
}

func (c *controller) UpdateUsername(id string, data *userdto.UserUpdateDto) (*api.Response, int, error) {
	return c.UserService.UpdateUsername(id, data)
}

func (c *controller) FindUsers(id string) (*api.Response, int, error) {
	return c.UserService.FindUsers(id)
}

func (c *controller) UpdateUserFriends(id string, friends []string) (*api.Response, int, error) {
	return c.UserService.UpdateUserFriends(id, friends)
}
