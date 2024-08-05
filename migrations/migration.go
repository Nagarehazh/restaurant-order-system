package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"sort"
)

type Migration struct {
	ID        uint   `gorm:"primaryKey"`
	Version   string `gorm:"unique"`
	AppliedAt int64
}

var migrations = []struct {
	version string
	up      func(*gorm.DB) error
	down    func(*gorm.DB) error
}{}

func registerMigration(version string, up, down func(*gorm.DB) error) {
	migrations = append(migrations, struct {
		version string
		up      func(*gorm.DB) error
		down    func(*gorm.DB) error
	}{version, up, down})
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Migration{}); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].version < migrations[j].version
	})

	for _, m := range migrations {
		var executed Migration
		if err := db.Where("version = ?", m.version).First(&executed).Error; err == gorm.ErrRecordNotFound {
			fmt.Printf("Running migration: %s\n", m.version)
			if err := m.up(db); err != nil {
				return fmt.Errorf("failed to run migration %s: %w", m.version, err)
			}
			db.Create(&Migration{Version: m.version, AppliedAt: db.NowFunc().Unix()})
		}
	}

	return nil
}

func MigrateDown(db *gorm.DB, steps int) error {
	var executedMigrations []Migration
	if err := db.Order("version DESC").Limit(steps).Find(&executedMigrations).Error; err != nil {
		return fmt.Errorf("failed to fetch executed migrations: %w", err)
	}

	for _, executed := range executedMigrations {
		for _, m := range migrations {
			if m.version == executed.Version {
				fmt.Printf("Rolling back migration: %s\n", m.version)
				if err := m.down(db); err != nil {
					return fmt.Errorf("failed to rollback migration %s: %w", m.version, err)
				}
				if err := db.Delete(&executed).Error; err != nil {
					return fmt.Errorf("failed to delete migration record %s: %w", m.version, err)
				}
				break
			}
		}
	}

	return nil
}
