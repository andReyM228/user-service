package main

import (
	"database/sql"
	"fmt"
	"log"

	"user_service/handler"
	"user_service/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "user_service"
)

type App struct {
	userRepo    repository.Repository
	userHandler handler.Handler

	router *fiber.App
}

func main() {
	db, err := initDatabase()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewRepository(db)
	userHandler := handler.NewHandler(userRepo)

	app := App{
		userRepo:    userRepo,
		userHandler: userHandler,
	}

	app.initHTTP()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	a.router.Get("v1/user-service", a.userHandler.Get)
	a.router.Post("v1/user-service", a.userHandler.Create)
	a.router.Put("v1/user-service", a.userHandler.Update)
	a.router.Delete("v1/user-service", a.userHandler.Delete)

	_ = a.router.Listen(":3000")
}

func initDatabase() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return sqlx.NewDb(db, ""), nil
}
