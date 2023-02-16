package main

import (
	"database/sql"
	"fmt"
	"log"
	cars2 "user_service/handler/cars"
	users2 "user_service/handler/users"
	"user_service/repository/cars"
	"user_service/repository/users"

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
	userRepo    users.Repository
	userHandler users2.Handler
	carRepo     cars.Repository
	carHandler  cars2.Handler

	router *fiber.App
}

func main() {
	db, err := initDatabase()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := users.NewRepository(db)
	userHandler := users2.NewHandler(userRepo)

	carRepo := cars.NewRepository(db)
	carsHandler := cars2.NewHandler(carRepo)

	app := App{
		userRepo:    userRepo,
		userHandler: userHandler,
		carRepo:     carRepo,
		carHandler:  carsHandler,
	}

	app.initHTTP()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	a.router.Get("v1/user-service/user/:id", a.userHandler.Get)
	a.router.Post("v1/user-service/user", a.userHandler.Create)
	a.router.Put("v1/user-service/user", a.userHandler.Update)
	a.router.Delete("v1/user-service/user/:id", a.userHandler.Delete)

	a.router.Get("v1/car-service/car/:id", a.carHandler.Get)
	a.router.Post("v1/car-service/car", a.carHandler.Create)
	a.router.Put("v1/car-service/car", a.carHandler.Update)
	a.router.Delete("v1/car-service/car/:id", a.carHandler.Delete)

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
