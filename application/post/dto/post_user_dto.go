package postdto

import (
	"errors"
	"go-api-insta/libs/logger"

	"github.com/go-playground/validator/v10"
)

type PostUserDto struct {
	UserId string `json:"userId" validate:"required"`
}

func (d *PostUserDto) Validate() error {

	validate := validator.New()

	err := validate.Struct(d)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logger.Development.Info(err.Error())
		}
		for _, e := range err.(validator.ValidationErrors) {
			err = errors.New(e.Field() + " " + e.Tag())
		}
	}

	return err
}
