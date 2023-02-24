package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	cars_handler "user_service/internal/handler/cars"
	users_handler "user_service/internal/handler/users"
	"user_service/internal/repository/cars"
	"user_service/internal/repository/users"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "user_service"
)

type App struct {
	serviceName string
	userRepo    users.Repository
	userHandler users_handler.Handler
	carRepo     cars.Repository
	carHandler  cars_handler.Handler
	logger      *logrus.Logger
	db          *sqlx.DB

	router *fiber.App
}

func New(name string) App {
	return App{
		serviceName: name,
	}
}

func (a *App) Run() {
	a.initLogger()
	a.initDatabase()
	a.initRepos()
	a.initHandlers()
	a.initHTTP()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	a.router.Get("v1/user-service/user/:id", a.userHandler.Get)
	a.router.Post("v1/user-service/user", a.userHandler.Create)
	a.router.Put("v1/user-service/user", a.userHandler.Update)
	a.router.Delete("v1/user-service/user/:id", a.userHandler.Delete)

	a.router.Get("v1/user-service/car/:id", a.carHandler.Get)
	a.router.Post("v1/user-service/car", a.carHandler.Create)
	a.router.Put("v1/user-service/car", a.carHandler.Update)
	a.router.Delete("v1/user-service/car/:id", a.carHandler.Delete)

	a.logger.Debug("fiber api started")
	_ = a.router.Listen(":3000")
}

func (a *App) initDatabase() {
	a.logger.Debug("opening database connection")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	a.db = db
	a.logger.Debug("database connected")
}

func (a *App) initLogger() {
	a.logger = logrus.New()
	a.logger.SetLevel(logrus.DebugLevel)
}

func (a *App) initRepos() {
	a.userRepo = users.NewRepository(a.db, a.logger)
	a.carRepo = cars.NewRepository(a.db, a.logger)
	a.logger.Debug("repos created")
}

func (a *App) initHandlers() {
	a.userHandler = users_handler.NewHandler(a.userRepo)
	a.carHandler = cars_handler.NewHandler(a.carRepo)
	a.logger.Debug("handlers created")
}
