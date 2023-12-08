package handlers

import (
	"errors"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"strconv"
	"swagger/restapi/operations"
	"swagger/services"
)

type (
	CreateUser struct {
		log      *services.Log
		sessions *services.Sessions
		users    *services.Users
	}
	ListUsers struct {
		log      *services.Log
		sessions *services.Sessions
		users    *services.Users
	}
)

func NewCreateUser(l *services.Log, s *services.Sessions, u *services.Users) CreateUser {
	return CreateUser{log: l, sessions: s, users: u}
}

func NewListUser(l *services.Log, s *services.Sessions, u *services.Users) ListUsers {
	return ListUsers{log: l, sessions: s, users: u}
}

func (c CreateUser) Handle(params operations.CreateUserParams) middleware.Responder {
	log := c.log.Func("createUser")
	switch {
	case params.Data == nil:
		log.BadRequest("data is null")
		return operations.NewCreateUserBadRequest()
	case params.Data.Name == nil:
		log.BadRequest("data.name is null")
		return operations.NewCreateUserBadRequest()
	case params.Data.Password == nil:
		log.BadRequest("data.password is null")
		return operations.NewCreateUserBadRequest()
	}
	id, fail := uuid.Parse(params.Session)
	if fail != nil {
		log.BadRequest("parse session id: %v", params.Session)
		return operations.NewCreateUserBadRequest()
	}
	session := c.sessions.Get(id)
	if session != nil {
		log.BadRequest("session not found: %v", id)
		return operations.NewCreateUserBadRequest()
	}
	fmt.Println("createUser from", session.User().Name())
	row, fail := c.users.New(*params.Data.Name, *params.Data.Password)
	switch {
	case fail == nil:
		log.OK(strconv.FormatUint(row.ID(), 10))
		return operations.NewCreateUserOK().WithPayload(row.ID())
	case errors.Is(fail, services.ErrUserIdExist):
		log.NotFound(fail.Error())
		return operations.NewCreateUserNotFound()
	case errors.Is(fail, services.ErrUserNameExist):
		log.NotFound(fail.Error())
		return operations.NewCreateUserNotFound()
	}
	log.InternalSerer(fail.Error())
	return operations.NewCreateUserInternalServerError()
}

func (l ListUsers) Handle(p operations.ListUsersParams) middleware.Responder {
	log := l.log.Func("listUsers")
	if p.Count == nil {
		log.BadRequest("count is null")
		return operations.NewListUsersBadRequest()
	}
	if p.Skip == nil {
		log.BadRequest("skip is null ")
		return operations.NewListUsersBadRequest()
	}
	list, fail := l.users.List(*p.Skip, *p.Count)
	if fail != nil {
		log.InternalSerer(fail.Error())
		return operations.NewListUsersInternalServerError()
	}
	payload := make([]*operations.ListUsersOKBodyItems0, len(list))
	for index, item := range list {
		payload[index] = &operations.ListUsersOKBodyItems0{ID: item.ID(), Name: item.Name()}
	}
	log.OK(strconv.Itoa(len(list)))
	return operations.NewListUsersOK().WithPayload(payload)
}
