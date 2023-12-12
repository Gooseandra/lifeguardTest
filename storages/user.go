package storages

import (
	"errors"
	"os"
	"strconv"
)

const (
	errUserLogin   = "login"
	errUserUnknown = "unknown user"
)

type (
	User interface {
		ID() UserID
		Name() UserName
		Surname() UserSurname
		Patronymic() UserPatronymic
		Email() UserEmail
		Vk() UserVk
		Tg() UserTg
		Nick() UserNick
		Password() UserPassword
		Phone() UserPhone
	}

	UserById map[UserID]User

	UserID = uint64

	UserIdExistError UserID

	UserName = string

	UserSurname = string

	UserPatronymic = string

	UserEmail = string

	UserVk = string

	UserTg = string

	UserNick = string

	UserNameExistError UserName

	UserNameMissingError UserName

	UserPassword = string

	UserPhone = string

	Users interface {
		ByName(name UserName) (User, error)
		ByID(id UserID) (User, error)
		New(name UserName, surname UserSurname, patronymic UserPatronymic, email UserEmail, vk UserVk, tg UserTg,
			nick UserNick, password UserPassword, phone UserPhone) (User, error)
		List(skip uint64, count uint32) ([]User, error)
	}
)

var (
	ErrUserExist    = os.ErrExist
	ErrUserLogin    = errors.New(errUserLogin)
	ErrUserNotFound = os.ErrNotExist
	ErrUserUnknown  = errors.New(errUserUnknown)
)

func (e UserIdExistError) Error() string {
	return "User already exist: id=" + strconv.FormatUint(UserID(e) /* !!! */, 10)
}

func (e UserNameExistError) Error() string {
	return "User already exist: name=" + string(e)
}

func (e UserNameMissingError) Error() string {
	return "User missing: name=" + string(e)
}
