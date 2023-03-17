package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r Repository) Get(id int64) error {

	return nil
}

func (r Repository) Update() error {

	return nil
}

func (r Repository) Create() error {

	return nil
}

func (r Repository) Delete(id int64) error {

	return nil
}