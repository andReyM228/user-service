package users

import (
	"errors"
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

func NewService(users users.Repository, log *logrus.Logger) Service {
	return Service{
		users: users,
		log:   log,
	}
}

func (s Service) Login(chatID int64, password string) error {

	return nil
}

func (s Service) Registration(user domain.User) error {
	_, err := s.users.Get(domain.FieldChatID, user.ChatID)
	if err == nil {
		return errors.New("this user already registered")
	}

	_, err = s.users.Get(domain.FieldPhone, user.Phone)
	if err == nil {
		return errors.New("this phone number already taken")
	}

	err = s.users.Create(user)
	if err != nil {
		return err
	}

	return nil
}

/*
для регистрации
1. сначала юзер нажимает /start, после этого ему выдаётся поле для ввода своих данных
2. мы должны проверить существует ли такой юзер. Если юзер существует, то вернуть сообщение "такой юзер уже существует", если не существует то создаём этого юзера
*/

