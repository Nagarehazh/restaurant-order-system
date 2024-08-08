package menu

import "gorm.io/gorm"

type Repository interface {
	CreateItem(item *Item) error
	GetItem(id uint) (*Item, error)
	UpdateItem(item *Item) error
	DeleteItem(id uint) error
	ListItems() ([]Item, error)
}

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) Repository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) CreateItem(item *Item) error {
	return r.DB.Create(item).Error
}

func (r *PostgresRepository) GetItem(id uint) (*Item, error) {
	var item Item
	if err := r.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *PostgresRepository) UpdateItem(item *Item) error {
	return r.DB.Save(item).Error
}

func (r *PostgresRepository) DeleteItem(id uint) error {
	return r.DB.Delete(&Item{}, id).Error
}

func (r *PostgresRepository) ListItems() ([]Item, error) {
	var items []Item
	if err := r.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
