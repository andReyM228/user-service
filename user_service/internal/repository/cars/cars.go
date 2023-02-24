package cars

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"user_service/internal/domain"
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

func (r Repository) Get(id int64) (domain.Car, error) {
	var car domain.Car

	if err := r.db.Get(&car, "SELECT * FROM cars WHERE id = $1", id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Infoln(err)
			return domain.Car{}, repository.NotFound{NotFound: "car"}
		}

		r.log.Errorln(err)
		return domain.Car{}, repository.InternalServerError{}
	}

	return car, nil
}

func (r Repository) Update(car domain.Car) error {
	_, err := r.db.Exec("UPDATE cars SET name = $1, model = $2 WHERE id = $3", car.Name, car.Model, car.ID)

	if err != nil {
		r.log.Errorln(err)
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Create(car domain.Car) error {
	if _, err := r.db.Exec("INSERT INTO cars (name, model) VALUES ($1, $2)", car.Name, car.Model); err != nil {
		r.log.Errorln(err)
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		r.log.Errorln(err)
		return repository.InternalServerError{}
	}

	return nil
}
