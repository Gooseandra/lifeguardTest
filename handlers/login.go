package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"swagger/restapi/operations"
	"swagger/services"
)

const errLoginUnknown = "unknown login error"

type Login struct {
	log      *services.Log
	sessions *services.Sessions
	users    *services.Users
}

var ErrUserUnknown = errors.New(errLoginUnknown)

func NewLogin(l *services.Log, s *services.Sessions, u *services.Users) Login {
	return Login{log: l, sessions: s, users: u}
}

func (l Login) Handle(p operations.LoginParams) middleware.Responder {
	var entity *services.SessionEntity
	log := l.log.Func("login")
	user, fail := l.users.ByName(p.Body.Name)
	if fail != nil {
		log.InternalSerer(fail.Error())
		return operations.NewListUsersInternalServerError()
	}
	if p.Body.Password != user.Password() {
		log.InternalSerer("invalid password")
		return operations.NewListUsersInternalServerError()
	}
	entity, fail = l.sessions.New(user)
	if fail != nil {
		log.InternalSerer(fail.Error())
		return operations.NewListUsersInternalServerError()
	}
	return operations.NewLoginOK().WithPayload(entity.ID().String())
}
