package postmodule

import (
	postdto "go-api-insta/application/post/dto"

	"go-api-insta/models/api"
)

// Controller defines the methods to be exposed in Controller layer
type Controller interface {
	CreatePost(userId string, data *postdto.PostDto) (*api.Response, int, error)
	FindAllPostUser(data *postdto.PostUserDto) (*api.Response, int, error)
	FindAllPostFriends(id string) (*api.Response, int, error)
}

type controller struct {
	PostService
}

func newController(service PostService) Controller {
	return &controller{
		PostService: service,
	}
}

func (c *controller) CreatePost(userId string, data *postdto.PostDto) (*api.Response, int, error) {
	return c.PostService.CreatePost(userId, data)
}

func (c *controller) FindAllPostUser(data *postdto.PostUserDto) (*api.Response, int, error) {
	return c.PostService.FindAllPostUser(data)
}

func (c *controller) FindAllPostFriends(id string) (*api.Response, int, error) {
	return c.PostService.FindAllPostFriends(id)
}
