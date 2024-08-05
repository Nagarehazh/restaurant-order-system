package menu

import "gorm.io/gorm"

type Repository interface {
	CreateMenuItem(item *MenuItem) error
	GetMenuItem(id uint) (*MenuItem, error)
	UpdateMenuItem(item *MenuItem) error
	DeleteMenuItem(id uint) error
	ListMenuItems() ([]MenuItem, error)
}

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) CreateMenuItem(item *MenuItem) error {
	return r.DB.Create(item).Error
}

func (r *PostgresRepository) GetMenuItem(id uint) (*MenuItem, error) {
	var item MenuItem
	if err := r.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *PostgresRepository) UpdateMenuItem(item *MenuItem) error {
	return r.DB.Save(item).Error
}

func (r *PostgresRepository) DeleteMenuItem(id uint) error {
	return r.DB.Delete(&MenuItem{}, id).Error
}

func (r *PostgresRepository) ListMenuItems() ([]MenuItem, error) {
	var items []MenuItem
	if err := r.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
