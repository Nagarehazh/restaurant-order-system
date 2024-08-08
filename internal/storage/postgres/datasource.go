package postgres

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Run(
	host string,
	user string,
	pass string,
	dbname string,
	port string,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=UTC",
		host, user, pass, dbname, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed connection")
	}

	log.Println("Database connected successfully")
	return db, nil
}
