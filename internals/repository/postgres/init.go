package postgres

import "github.com/BigNutJaa/users/internals/infrastructure/database"

type PostgresRepository struct {
	db *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}
