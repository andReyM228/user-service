package users

import (
	"errors"

	"user_service/internal/domain"
	"user_service/internal/domain/errs"
	"user_service/internal/repository"
	"user_service/internal/repository/cars"
	"user_service/internal/repository/transfers"
	"user_service/internal/repository/user_cars"
	"user_service/internal/repository/users"

	"github.com/andReyM228/lib/log"
)

const systemUser = 0

type Service struct {
	users     users.Repository
	cars      cars.Repository
	userCars  user_cars.Repository
	transfers transfers.Repository
	log       log.Logger
}

func NewService(users users.Repository, log log.Logger) Service {
	return Service{
		users: users,
		log:   log,
	}
}

func (s Service) Login(chatID int64, password string) (int64, error) {
	user, err := s.users.Get(domain.FieldChatID, chatID)
	if err != nil {
		if errors.As(err, &repository.NotFound{}) {
			return 0, errs.NotFoundError{What: "user"}
		}

		s.log.Error(err.Error())

		return 0, errs.InternalError{}
	}

	if password != user.Password {
		return 0, errs.Unauthorized{Cause: "wrong password"}
	}

	return int64(user.ID), nil
}

/*
для логина
1. попробовать получить пользователя по chatID.
2. если его нет, то возвращаем ошибку not found.
3. если он есть и мы его получили, то сравниваем пароль в бд с паролем который ввёл пользователь.
4. если пароли не совпали, то возвращаем ошибку unauthorized.
*/
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
