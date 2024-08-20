package repositories

import (
	"database/sql"
	"errors"
	"test-backend-pari/items"
	"test-backend-pari/items/entities"
	"time"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) items.Repository {
	return &ItemRepository{
		db: db,
	}
}

// Create implements items.Repository.
func (repo *ItemRepository) Create(item entities.CreateItemRequest) error {
	_, err := repo.db.Query(
		"INSERT INTO item (name, created_on, updated_on) VALUES ($1, $2, $3)",
		item.Name,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements items.Repository.
func (repo *ItemRepository) Delete(id int) error {
	_, err := repo.db.Query("DELETE FROM item WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements items.Repository.
func (repo *ItemRepository) GetAll() ([]entities.Item, error) {
	var items []entities.Item

	rows, err := repo.db.Query("SELECT id, name, created_on, updated_on FROM item ORDER BY created_on DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entities.Item
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.CreatedOn,
			&item.UpdatedOn,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// GetById implements items.Repository.
func (repo *ItemRepository) GetById(id int) (entities.Item, error) {
	var item entities.Item
	err := repo.db.QueryRow("SELECT id, name, created_on, updated_on FROM item WHERE id = $1", id).Scan(
		&item.Id,
		&item.Name,
		&item.CreatedOn,
		&item.UpdatedOn,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Item{}, errors.New("item not found")
		}
		return entities.Item{}, err
	}

	return item, nil
}

// Update implements items.Repository.
func (repo *ItemRepository) Update(item entities.Item) error {
	_, err := repo.db.Query(
		"UPDATE item SET name = $2, updated_on = $3 WHERE id = $1",
		item.Id,
		item.Name,
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
