package migrations

import (
	"gorm.io/gorm"
)

func init() {
	registerMigration("20230805000000", createOrdersTables, dropOrdersTables)
}

func createOrdersTables(db *gorm.DB) error {
	return db.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id SERIAL PRIMARY KEY,
            created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC'),
            updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC'),
            deleted_at TIMESTAMP WITHOUT TIME ZONE,
            user_id INTEGER NOT NULL,
            status VARCHAR(20) NOT NULL DEFAULT 'pending',
            total DECIMAL(10, 2) NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );

        CREATE TABLE IF NOT EXISTS order_items (
            id SERIAL PRIMARY KEY,
            created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC'),
            updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC'),
            deleted_at TIMESTAMP WITHOUT TIME ZONE,
            order_id INTEGER NOT NULL,
            menu_item_id INTEGER NOT NULL,
            quantity INTEGER NOT NULL,
            price DECIMAL(10, 2) NOT NULL,
            FOREIGN KEY (order_id) REFERENCES orders(id),
            FOREIGN KEY (menu_item_id) REFERENCES menu_items(id)
        );
    `).Error
}

func dropOrdersTables(db *gorm.DB) error {
	return db.Exec(`
        DROP TABLE IF EXISTS order_items;
        DROP TABLE IF EXISTS orders;
    `).Error
}
