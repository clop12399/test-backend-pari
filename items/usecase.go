package items

import "test-backend-pari/items/entities"

type Usecase interface {
	GetAll() ([]entities.Item, error)
	GetById(id int) (entities.Item, error)
	Create(payload entities.CreateItemRequest) error
	Update(payload entities.UpdateItemRequest, id int) error
	Delete(id int) error
}