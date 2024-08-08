package order

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Create(order *Order) error
	GetById(id uint) (*Order, error)
	Update(order *Order) error
	Delete(id uint) error
	List(userID uint) ([]Order, error)
}

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) Repository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) Create(order *Order) error {
	return r.DB.Create(order).Error
}

func (r *PostgresRepository) GetById(id uint) (*Order, error) {
	var order Order
	if err := r.DB.Preload("Items.MenuItem").First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

func (r *PostgresRepository) Update(order *Order) error {
	return r.DB.Save(order).Error
}

func (r *PostgresRepository) Delete(id uint) error {
	return r.DB.Delete(&Order{}, id).Error
}

func (r *PostgresRepository) List(userID uint) ([]Order, error) {
	var orders []Order
	query := r.DB.Preload("Items").Preload("Items.MenuItem")

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
