package order

import (
	"gorm.io/gorm"
	"restaurant-order-system/internal/auth"
	"restaurant-order-system/internal/menu"
)

type Status string

const (
	OrderStatusPending   Status = "pending"
	OrderStatusPreparing Status = "preparing"
	OrderStatusReady     Status = "ready"
	OrderStatusDelivered Status = "delivered"
	OrderStatusCancelled Status = "cancelled"
)

type Order struct {
	gorm.Model
	UserID uint      `json:"user_id" gorm:"not null"`
	User   auth.User `json:"-" gorm:"foreignKey:UserID"`
	Status Status    `json:"status" gorm:"type:varchar(20);not null;default:'pending'"`
	Total  float64   `json:"total" gorm:"type:decimal(10,2);not null"`
	Items  []Item    `json:"items" gorm:"foreignKey:OrderID"`
}

type Item struct {
	gorm.Model
	OrderID    uint      `json:"order_id" gorm:"not null"`
	MenuItemID uint      `json:"menu_item_id" gorm:"not null"`
	MenuItem   menu.Item `json:"menu_item" gorm:"foreignKey:MenuItemID"`
	Quantity   int       `json:"quantity" gorm:"not null"`
	Price      float64   `json:"price" gorm:"type:decimal(10,2);not null"`
}

func (Item) TableName() string {
	return "order_items"
}
