package storages

import (
	"errors"
	"os"
	"strconv"
	"time"
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
		Apply() *UserTime
		Fired() *UserTime
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

	UserTime = time.Time

	Users interface {
		ByName(name UserName) (User, error)
		ByID(id UserID) (User, error)
		New(name UserName, surname UserSurname, patronymic UserPatronymic, email UserEmail, vk UserVk, tg UserTg,
			nick UserNick, password UserPassword, phone UserPhone, apply *UserTime) (User, error)
		List(skip uint64, count uint32) ([]User, error)
		Update(id UserID, name UserName, surname UserSurname, patronymic UserPatronymic, email UserEmail, vk UserVk,
			tg UserTg, nick UserNick, phone UserPhone, apply *UserTime) (User, error)
		Fired(id UserID, time UserTime) (User, error)
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
