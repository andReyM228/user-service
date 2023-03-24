package app

import (
	"fmt"
	"log"
	"net/http"
	"user_service/internal/config"
	car_trading_handler "user_service/internal/handler/car_trading"
	"user_service/internal/repository/transfers"
	"user_service/internal/repository/user_cars"
	"user_service/internal/service/car_trading"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	cars_handler "user_service/internal/handler/cars"
	users_handler "user_service/internal/handler/users"
	"user_service/internal/repository/cars"
	"user_service/internal/repository/users"
)

type App struct {
	config            config.Config
	serviceName       string
	userRepo          users.Repository
	userHandler       users_handler.Handler
	carRepo           cars.Repository
	carHandler        cars_handler.Handler
	carTradingService car_trading.Service
	userCarsRepo      user_cars.Repository
	transferRepo      transfers.Repository
	carTradingHandler car_trading_handler.Handler
	logger            *logrus.Logger
	db                *sqlx.DB
	clientHTTP        *http.Client

	router *fiber.App
}

func New(name string) App {
	return App{
		serviceName: name,
	}
}

func (a *App) Run() {
	a.populateConfig()
	a.initLogger()
	a.initDatabase()
	a.initHTTPClient()
	a.initRepos()
	a.initServices()
	a.initHandlers()
	a.initHTTP()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	a.router.Post("v1/user-service/buy-car/:user_id/:car_id", a.carTradingHandler.BuyCar)

	a.router.Get("v1/user-service/user/:id", a.userHandler.Get)
	a.router.Post("v1/user-service/user", a.userHandler.Create)
	a.router.Put("v1/user-service/user", a.userHandler.Update)
	a.router.Delete("v1/user-service/user/:id", a.userHandler.Delete)

	a.router.Get("v1/user-service/car/:id", a.carHandler.Get)
	a.router.Post("v1/user-service/car", a.carHandler.Create)
	a.router.Put("v1/user-service/car", a.carHandler.Update)
	a.router.Delete("v1/user-service/car/:id", a.carHandler.Delete)

	a.logger.Debug("fiber api started")
	_ = a.router.Listen(fmt.Sprintf(":%d", a.config.HTTP.Port))
}

func (a *App) initDatabase() {
	a.logger.Debug("opening database connection")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		a.config.DB.Host, a.config.DB.Port, a.config.DB.User, a.config.DB.Password, a.config.DB.DBname)

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
	a.userCarsRepo = user_cars.NewRepository(a.db, a.logger)
	a.userRepo = users.NewRepository(a.db, a.logger)
	a.carRepo = cars.NewRepository(a.db, a.logger)
	a.transferRepo = transfers.NewRepository(a.clientHTTP, a.logger)
	a.logger.Debug("repos created")
}

func (a *App) initHandlers() {
	a.userHandler = users_handler.NewHandler(a.userRepo)
	a.carHandler = cars_handler.NewHandler(a.carRepo)
	a.carTradingHandler = car_trading_handler.NewHandler(a.carTradingService)
	a.logger.Debug("handlers created")
}

func (a *App) initServices() {
	a.carTradingService = car_trading.NewService(a.userRepo, a.carRepo, a.userCarsRepo, a.transferRepo, a.logger)

	a.logger.Debug("services created")
}

func (a *App) populateConfig() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal()
	}

	a.config = cfg
}

func (a *App) initHTTPClient() {
	a.clientHTTP = http.DefaultClient
}
