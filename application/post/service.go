package postmodule

import (
	postdto "go-api-insta/application/post/dto"
	"go-api-insta/libs/logger"
	"go-api-insta/models/api"
	"go-api-insta/models/post"
	"go-api-insta/models/user"

	"github.com/gofiber/fiber/v2"
)

type PostService interface {
	CreatePost(userId string, data *postdto.PostDto) (*api.Response, int, error)
	FindAllPostUser(data *postdto.PostUserDto) (*api.Response, int, error)
	FindAllPostFriends(id string) (*api.Response, int, error)
}

type postService struct {
	postRepository post.PostRepository
	userRepository user.UserRepository
}

func newService() PostService {
	return &postService{}
}

func (p *postService) CreatePost(userId string, data *postdto.PostDto) (*api.Response, int, error) {

	newPost := p.postDtoToPost(userId, data)

	_, err := p.postRepository.CreatePost(userId, newPost)
	if err != nil {
		logger.Production.Info(err.Error())
		return &api.Response{Error: true, ErrorMessage: err.Error(), Payload: "error insert post"}, fiber.StatusInternalServerError, err
	}
	return &api.Response{Error: false, Payload: "success insert post"}, fiber.StatusOK, nil
}

func (p *postService) FindAllPostUser(data *postdto.PostUserDto) (*api.Response, int, error) {
	postsUser, err := p.postRepository.FindAllPostByUser(data.UserId)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error(), Payload: "error insert user"}, fiber.StatusBadRequest, err
	}
	return &api.Response{Error: false, Payload: postsUser}, fiber.StatusOK, nil
}

func (p *postService) FindAllPostFriends(id string) (*api.Response, int, error) {

	user, err := p.userRepository.GetUserById(id)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error(), Payload: "error insert user"}, fiber.StatusBadRequest, err
	}

	friendAndUser := append(user.Friends, id)

	postsUser, err := p.postRepository.FindAllPostFriends(friendAndUser)
	if err != nil {
		return &api.Response{Error: true, ErrorMessage: err.Error(), Payload: "error insert user"}, fiber.StatusBadRequest, err
	}
	return &api.Response{Error: false, Payload: postsUser}, fiber.StatusOK, nil
}

func (p *postService) postDtoToPost(id string, postDto *postdto.PostDto) *post.Post {
	return &post.Post{Description: postDto.Description, File: postDto.File, UserId: id}
}
