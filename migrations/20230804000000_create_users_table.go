package migrations

import (
	"gorm.io/gorm"
)

func init() {
	registerMigration("20230804000000", createUsersTable, dropUsersTable)
}

func createUsersTable(db *gorm.DB) error {
	return db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC'),
            updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC'),
            deleted_at TIMESTAMP WITHOUT TIME ZONE,
            username VARCHAR(255) NOT NULL UNIQUE,
            email VARCHAR(255) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL
        )
    `).Error
}

func dropUsersTable(db *gorm.DB) error {
	return db.Exec("DROP TABLE IF EXISTS users").Error
}
