package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Item struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type CreateItemRequest struct {
	Name string `json:"name" validate:"required"`
}

func (request CreateItemRequest) Validate() error {
	return validator.New().Struct(request)
}

type UpdateItemRequest struct {
	Name string `json:"name" validate:"required"`
}

func (request UpdateItemRequest) Validate() error {
	return validator.New().Struct(request)
}
