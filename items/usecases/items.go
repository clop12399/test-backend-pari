package usecases

import (
	"test-backend-pari/items"
	"test-backend-pari/items/entities"
)

type ItemUsecase struct {
	repo items.Repository
}

func NewItemUsecase(repo items.Repository) items.Usecase {
	return &ItemUsecase{
		repo: repo,
	}
}

// Create implements items.Usecase.
func (uc *ItemUsecase) Create(payload entities.CreateItemRequest) error {
	return uc.repo.Create(payload)
}

// Delete implements items.Usecase.
func (uc *ItemUsecase) Delete(id int) error {
	_, err := uc.repo.GetById(id)
	if err != nil {
		return err
	}

	return uc.repo.Delete(id)
}

// GetAll implements items.Usecase.
func (uc *ItemUsecase) GetAll() ([]entities.Item, error) {
	return uc.repo.GetAll()
}

// GetById implements items.Usecase.
func (uc *ItemUsecase) GetById(id int) (entities.Item, error) {
	return uc.repo.GetById(id)
}

// Update implements items.Usecase.
func (uc *ItemUsecase) Update(payload entities.UpdateItemRequest, id int) error {
	_, err := uc.repo.GetById(id)
	if err != nil {
		return err
	}

	return uc.repo.Update(entities.Item{
		Id:   id,
		Name: payload.Name,
	})
}
