package menu

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

func (Item) TableName() string {
	return "menu_items"
}
