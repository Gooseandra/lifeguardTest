package services

import (
	"errors"
	"fmt"
	"strconv"
	"swagger/storages"
)

const (
	errUserIdExist   = "user id exist"
	errUserNameExist = "user name exist"
)

type (
	UserIdExistError   storages.UserID
	UserNameExistError storages.UserName
	Users              struct{ storage storages.Users }
)

var (
	ErrUserIdExist   = errors.New(errUserIdExist)
	ErrUserNameExist = errors.New(errUserNameExist)
)

func NewUsers(s storages.Users) *Users { return &Users{storage: s} }

func (e UserIdExistError) Error() string {
	return e.Unwrap().Error() + ": " + strconv.FormatUint(storages.UserID(e), 10)
}

func (e UserIdExistError) Unwrap() error { return ErrUserIdExist }

func (e UserNameExistError) Error() string {
	return e.Unwrap().Error() + ": " + storages.UserName(e)
}

func (e UserNameExistError) Unwrap() error { return ErrUserNameExist }

func (u Users) ByName(n storages.UserName) (storages.User, error) {
	row, fail := u.storage.ByName(n)
	switch fail {
	case storages.UserNameMissingError(n):
	}
	return row, nil
}

func (u Users) ByID(id storages.UserID) (storages.User, error) {
	users, err := u.storage.ByID(id)
	if err == nil {
		return users, nil
	}
	return nil, fmt.Errorf("user serice: %w", err)
}

func (u Users) List(s uint64, c uint32) ([]storages.User, error) {
	users, err := u.storage.List(s, c)
	if err == nil {
		return users, nil
	}
	return nil, fmt.Errorf("user serice: %w", err)
}

func (u Users) New(name storages.UserName, surname storages.UserSurname, patronymic storages.UserPatronymic,
	email storages.UserEmail, vk storages.UserVk, tg storages.UserTg, nick storages.UserNick,
	password storages.UserPassword, phone storages.UserPhone) (storages.User, error) {

	var idExistError storages.UserIdExistError
	var nameExistError storages.UserNameExistError
	user, err := u.storage.New(name, surname, patronymic, email, vk, tg, nick, password, phone)
	if err == nil {
		return user, nil
	}
	switch {
	case errors.As(err, &idExistError):
		return nil, UserIdExistError(idExistError)
	case errors.As(err, &nameExistError):
		return nil, UserNameExistError(nameExistError)
	}
	return nil, fmt.Errorf("user serice: %w", err)
}
