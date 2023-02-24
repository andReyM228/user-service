package users

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"user_service/repository"

	"user_service/domain"
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

func (r Repository) Get(id int64) (domain.User, error) {
	var user domain.User

	if err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Infoln(err)
			return domain.User{}, repository.NotFound{NotFound: "user"}
		}

		r.log.Errorln(err)
		return domain.User{}, repository.InternalServerError{}
	}

	return user, nil
}

func (r Repository) Update(user domain.User) error {
	if _, err := r.db.Exec("UPDATE users SET name = $1, surname = $2, phone = $3, email = $4 WHERE id = $5",
		user.Name, user.Surname, user.Phone, user.Email, user.ID); err != nil {
		r.log.Errorln(err)
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Create(user domain.User) error {
	if _, err := r.db.Exec("INSERT INTO users (name, surname, phone, email) VALUES ($1, $2, $3, $4)", user.Name, user.Surname, user.Phone, user.Email); err != nil {
		r.log.Errorln(err)
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	if _, err := r.db.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		r.log.Errorln(err)
		return repository.InternalServerError{}
	}

	return nil
}
