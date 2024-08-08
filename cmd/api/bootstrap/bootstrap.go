package bootstrap

import (
	"fmt"
	"restaurant-order-system/internal/server"
	"restaurant-order-system/internal/shared/enviroments"
	"restaurant-order-system/internal/storage/postgres"
	"restaurant-order-system/internal/storage/postgres/migrations"
)

func Run() error {
	db, err := postgres.Run(
		enviroments.PostgresHost,
		enviroments.PostgresUser,
		enviroments.PostgresPass,
		enviroments.PostgresDBName,
		enviroments.PostgresPort,
	)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = migrations.Run(db)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	srv := server.New(enviroments.Host, enviroments.Port, db)
	return srv.Run()
}
