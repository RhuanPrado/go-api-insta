package authmodule

import (
	authdto "go-api-insta/application/auth/dto"
	"go-api-insta/models/api"
)

type Controller interface {
	Authorization(data *authdto.AuthDto) (*api.Response, int, error)
}

type controller struct {
	AuthService
}

func newController(service AuthService) Controller {
	return &controller{
		AuthService: service,
	}
}

func (c *controller) Authorization(data *authdto.AuthDto) (*api.Response, int, error) {
	return c.AuthService.Authorization(data)
}
