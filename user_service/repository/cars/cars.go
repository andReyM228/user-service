package cars

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(database *sqlx.DB) Repository {
	return Repository{
		db: database,
	}
}
