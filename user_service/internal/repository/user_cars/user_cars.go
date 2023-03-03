package user_cars

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"user_service/internal/repository"
)

type Repository struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewRepository(database *sqlx.DB, log *logrus.Logger) Repository {
	return Repository{
		db:  database,
		log: log,
	}
}

func (r Repository) Create(userID, carID int) error {
	if _, err := r.db.Exec("INSERT INTO user_cars (user_id, car_id) VALUES ($1, $2)", userID, carID); err != nil {
		r.log.Errorln(err)
		return repository.InternalServerError{}
	}

	return nil
}
