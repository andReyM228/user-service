package car_trading

import (
	"github.com/sirupsen/logrus"
	"user_service/internal/domain"
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
	log       *logrus.Logger
}

func NewService(users users.Repository, cars cars.Repository, userCars user_cars.Repository, transfers transfers.Repository, log *logrus.Logger) Service {
	return Service{
		users:     users,
		cars:      cars,
		userCars:  userCars,
		transfers: transfers,
		log:       log,
	}
}

func (s Service) BuyCar(userID, carID int64) error {
	user, err := s.users.Get(domain.FieldID, userID)
	if err != nil {
		s.log.Errorln(err)
		return err
	}

	car, err := s.cars.Get(carID)
	if err != nil {
		s.log.Errorln(err)
		return err
	}

	s.log.Infoln("Sending transfer")
	if err := s.transfers.Transfer(user.ID, systemUser, int(car.Price)); err != nil {
		s.log.Errorln(err)
		return err
	}
	s.log.Infoln("Transfer sent")

	if err := s.userCars.Create(user.ID, car.ID); err != nil {
		s.log.Errorln(err)
		return err
	}
	s.log.Infoln("Car sent")

	return nil
}

/*
1. принимаем: user_id - пользователь который хочет купить машину, car_id - машина которую хочет пользователь
2. проверяем существует ли пользователь и машина
3. делаем транзакцию
4. если транзакция успешна, добавляем пользователю машину, если не успешна, то выдаём ошибку
*/
