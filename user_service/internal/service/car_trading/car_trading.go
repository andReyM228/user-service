package car_trading

import (
	"errors"
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/log"
	"user_service/internal/domain"
	"user_service/internal/repository"
	"user_service/internal/repository/cars"
	"user_service/internal/repository/transfers"
	"user_service/internal/repository/user_cars"
	"user_service/internal/repository/users"
)

const systemUser = 0

type Service struct {
	users     users.Repository
	cars      cars.Repository
	userCars  user_cars.Repository
	transfers transfers.Repository
	log       log.Logger
}

func NewService(users users.Repository, cars cars.Repository, userCars user_cars.Repository, transfers transfers.Repository, log log.Logger) Service {
	return Service{
		users:     users,
		cars:      cars,
		userCars:  userCars,
		transfers: transfers,
		log:       log,
	}
}

func (s Service) BuyCar(chatID, carID int64) error {
	user, err := s.users.Get(domain.FieldChatID, chatID)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}

	car, err := s.cars.Get(carID)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}

	s.log.Info("Sending transfer")
	if err := s.transfers.Transfer(user.ID, systemUser, int(car.Price)); err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info("Transfer sent")

	if err := s.userCars.Create(user.ID, car.ID); err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info("Car sent")

	return nil
}

/*
1. принимаем: user_id - пользователь который хочет купить машину, car_id - машина которую хочет пользователь
2. проверяем существует ли пользователь и машина
3. делаем транзакцию
4. если транзакция успешна, добавляем пользователю машину, если не успешна, то выдаём ошибку
*/

func (s Service) GetCar(id int64) (domain.Car, error) {
	car, err := s.cars.Get(id)
	if err != nil {
		if errors.As(err, &repository.InternalServerError{}) {
			s.log.Error(err.Error())
			return domain.Car{}, errs.InternalError{}
		}
		s.log.Debug(err.Error())

		return domain.Car{}, errs.NotFoundError{What: err.Error()}
	}

	return car, nil
}
