package app

import (
	"embed"
	"fmt"
	"github.com/andReyM228/lib/bus"
	"github.com/andReyM228/lib/database"
	"github.com/andReyM228/lib/rabbit"
	"github.com/gofiber/fiber/v2"
	"net/http"

	"user_service/internal/config"
	car_trading_handler "user_service/internal/handler/car_trading"
	cars_handler "user_service/internal/handler/cars"
	users_handler "user_service/internal/handler/users"
	"user_service/internal/repository/cars"
	"user_service/internal/repository/transfers"
	"user_service/internal/repository/user_cars"
	"user_service/internal/repository/users"
	"user_service/internal/service/car_trading"
	users_service "user_service/internal/service/users"

	"github.com/andReyM228/lib/log"
	_ "github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const urlRabbit = "amqp://guest:guest@rabbitmq:5672/"

type App struct {
	config            config.Config
	serviceName       string
	userRepo          users.Repository
	userHandler       users_handler.Handler
	carRepo           cars.Repository
	carHandler        cars_handler.Handler
	carTradingService car_trading.Service
	userService       users_service.Service
	userCarsRepo      user_cars.Repository
	transferRepo      transfers.Repository
	carTradingHandler car_trading_handler.Handler
	logger            log.Logger
	db                *sqlx.DB
	clientHTTP        *http.Client
	rabbit            rabbit.Rabbit

	router *fiber.App
}

func New(name string) App {
	return App{
		serviceName: name,
	}
}

func (a *App) Run(fs embed.FS) {
	a.populateConfig()
	a.initLogger()
	a.initDatabase(fs)
	a.initRabbit()
	a.initHTTPClient()
	a.initRepos()
	a.initServices()
	a.initHandlers()
	a.listenRabbit()
	a.initHTTP()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	a.router.Post("v1/user-service/buy-car/:chat_id/:car_id", a.carTradingHandler.BuyCar)
	a.router.Post("v1/user-service/sell-car/:chat_id/:car_id", a.carTradingHandler.SellCar)

	a.router.Get("v1/user-service/user/:id", a.userHandler.Get)
	a.router.Post("v1/user-service/user", a.userHandler.Create)
	a.router.Post("v1/user-service/user/login", a.userHandler.Login)
	a.router.Put("v1/user-service/user", a.userHandler.Update)
	a.router.Delete("v1/user-service/user/:id", a.userHandler.Delete)

	a.router.Get("v1/user-service/car/:id", a.carHandler.Get)
	a.router.Get("v1/user-service/cars/:label", a.carHandler.GetAll)
	a.router.Get("v1/user-service/user-cars", a.carHandler.GetUserCars)
	a.router.Post("v1/user-service/car", a.carHandler.Create)
	a.router.Put("v1/user-service/car", a.carHandler.Update)
	a.router.Delete("v1/user-service/car/:id", a.carHandler.Delete)

	a.logger.Debug("fiber api started")
	_ = a.router.Listen(fmt.Sprintf(":%d", a.config.HTTP.Port))
}

func (a *App) listenRabbit() {
	go func() {
		err := a.rabbit.Consume(bus.SubjectUserServiceCreateUser, a.userHandler.BrokerCreate)
		if err != nil {
			return
		}
	}()

	go func() {
		err := a.rabbit.Consume(bus.SubjectUserServiceLoginUser, a.userHandler.BrokerLogin)
		if err != nil {
			return
		}
	}()

	go func() {
		err := a.rabbit.Consume(bus.SubjectUserServiceGetUserByID, a.userHandler.BrokerGetUserByID)
		if err != nil {
			return
		}
	}()

	go func() {
		err := a.rabbit.Consume(bus.SubjectUserServiceGetCarByID, a.carHandler.BrokerGetCarByID)
		if err != nil {
			return
		}
	}()
}

func (a *App) initDatabase(fs embed.FS) {
	database.InitDatabase(a.logger, a.config.DB, fs)
}

func (a *App) initLogger() {
	a.logger = log.Init()
}

func (a *App) initRepos() {
	a.userCarsRepo = user_cars.NewRepository(a.db, a.logger)
	a.userRepo = users.NewRepository(a.db, a.logger)
	a.carRepo = cars.NewRepository(a.db, a.logger)
	a.transferRepo = transfers.NewRepository(a.clientHTTP, a.logger)
	a.logger.Debug("repos created")
}

func (a *App) initHandlers() {
	a.userHandler = users_handler.NewHandler(a.userRepo, a.userService, a.rabbit)
	a.carHandler = cars_handler.NewHandler(a.carRepo, a.carTradingService, a.rabbit)
	a.carTradingHandler = car_trading_handler.NewHandler(a.carTradingService)
	a.logger.Debug("handlers created")
}

func (a *App) initServices() {
	a.carTradingService = car_trading.NewService(a.userRepo, a.carRepo, a.userCarsRepo, a.transferRepo, a.logger)
	a.userService = users_service.NewService(a.userRepo, a.logger)

	a.logger.Debug("services created")
}

func (a *App) populateConfig() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Init().Fatal(err.Error())
	}

	a.config = cfg
}

func (a *App) initHTTPClient() {
	a.clientHTTP = http.DefaultClient
}

func (a *App) initRabbit() {
	var err error
	a.rabbit, err = rabbit.NewRabbitMQ(urlRabbit)
	if err != nil {
		a.logger.Fatal(err.Error())
	}
}
