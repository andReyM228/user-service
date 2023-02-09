package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(database *sqlx.DB) Repository {
	return Repository{
		db: database,
	}
}

func (r Repository) Get(ctx *fiber.Ctx) error {
	return nil
}

func (r Repository) Update(ctx *fiber.Ctx) error {
	return nil
}

func (r Repository) Create(ctx *fiber.Ctx) error {
	return nil
}

func (r Repository) Delete(ctx *fiber.Ctx) error {
	return nil
}
