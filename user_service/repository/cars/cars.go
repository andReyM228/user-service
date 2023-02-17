package cars

import (
	"github.com/jmoiron/sqlx"
	"user_service/domain"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(database *sqlx.DB) Repository {
	return Repository{
		db: database,
	}
}

func (r Repository) Get(id int64) (domain.Car, error) {
	var car domain.Car

	if err := r.db.Get(&car, "SELECT * FROM cars WHERE id = $1", id); err != nil {
		return domain.Car{}, nil
	}

	return car, nil
}

func (r Repository) Update(car domain.Car) error {
	if _, err := r.db.Exec("UPDATE cars SET name = $1, model = $2 WHERE id = $3",
		car.Name, car.Model, car.ID); err != nil {
		return err
	}

	return nil
}

func (r Repository) Create(car domain.Car) error {
	if _, err := r.db.Exec("INSERT INTO cars (name, model) VALUES ($1, $2)", car.Name, car.Model); err != nil {
		return err
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	if _, err := r.db.Exec("DELETE FROM cars WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}
