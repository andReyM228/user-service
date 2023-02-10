package repository

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

func (r Repository) Get(id int64) (domain.User, error) {
	var user domain.User

	if err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r Repository) Update(user domain.User) error {
	if _, err := r.db.Exec("UPDATE users SET name = $1, surname = $2, phone = $3, email = $4 WHERE id = $5",
		user.Name, user.Surname, user.Phone, user.Email, user.ID); err != nil {
		return err
	}

	return nil
}

func (r Repository) Create(user domain.User) error {
	if _, err := r.db.Exec("INSERT INTO users (name, surname, phone, email) VALUES ($1, $2, $3, $4)", user.Name, user.Surname, user.Phone, user.Email); err != nil {
		return err
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	if _, err := r.db.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}
