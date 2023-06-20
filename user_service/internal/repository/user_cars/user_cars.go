package user_cars

import (
	"user_service/internal/repository"

	"github.com/andReyM228/lib/log"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewRepository(database *sqlx.DB, log log.Logger) Repository {
	return Repository{
		db:  database,
		log: log,
	}
}

func (r Repository) Create(userID, carID int) error {
	if _, err := r.db.Exec("INSERT INTO user_cars (user_id, car_id) VALUES ($1, $2)", userID, carID); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}
