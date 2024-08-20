package items

import "test-backend-pari/items/entities"

type Repository interface {
	GetAll() ([]entities.Item, error)
	GetById(id int) (entities.Item, error)
	Create(item entities.CreateItemRequest) error
	Update(item entities.Item) error
	Delete(id int) error
}
