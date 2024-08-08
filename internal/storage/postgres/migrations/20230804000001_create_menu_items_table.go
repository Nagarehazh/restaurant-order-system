package migrations

import (
	"gorm.io/gorm"
)

func init() {
	registerMigration("20230804000001", createMenuItemsTable, dropMenuItemsTable)
}

func createMenuItemsTable(db *gorm.DB) error {
	return db.Exec(`
        CREATE TABLE IF NOT EXISTS menu_items (
            id SERIAL PRIMARY KEY,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP WITH TIME ZONE,
            name VARCHAR(255) NOT NULL,
            description TEXT,
            price DECIMAL(10, 2) NOT NULL,
            category VARCHAR(100)
        )
    `).Error
}

func dropMenuItemsTable(db *gorm.DB) error {
	return db.Exec("DROP TABLE IF EXISTS menu_items").Error
}
