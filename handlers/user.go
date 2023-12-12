package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"strconv"
	"swagger/restapi/operations"
	"swagger/services"
	"time"
)

type (
	user struct {
		log      *services.Log
		sessions *services.Sessions
		users    *services.Users
	}
	CreateUser struct{ user }
	GetUser    struct {
		log      *services.Log
		sessions *services.Sessions
		users    *services.Users
	}
	ListUsers struct {
		log      *services.Log
		sessions *services.Sessions
		users    *services.Users
	}
	UpdateUser struct{ user }
	FiredUser  struct{ user }
)

func NewCreateUser(l *services.Log, s *services.Sessions, u *services.Users) CreateUser {
	return CreateUser{user: user{log: l, sessions: s, users: u}}
}

func NewGetUser(l *services.Log, s *services.Sessions, u *services.Users) GetUser {
	return GetUser{log: l, sessions: s, users: u}
}

func NewListUser(l *services.Log, s *services.Sessions, u *services.Users) ListUsers {
	return ListUsers{log: l, sessions: s, users: u}
}

func NewUpdateUser(l *services.Log, s *services.Sessions, u *services.Users) UpdateUser {
	return UpdateUser{user: user{log: l, sessions: s, users: u}}
}

func NewFiredUser(l *services.Log, s *services.Sessions, u *services.Users) FiredUser {
	return FiredUser{user: user{log: l, sessions: s, users: u}}
}

func (f FiredUser) Handle(params operations.FiredUserParams) middleware.Responder {
	log := f.log.Func("firedUser")
	time, err := time.Parse(time.DateTime, *params.Body.Fired)
	if err != nil {
		log.BadRequest("Invalid time format")
		return operations.NewCreateCallBadRequest()
	}
	row, fail := f.users.Fired(params.ID, time)
	if fail != nil {
		log.NotFound(fail.Error())
		return operations.NewFiredUserInternalServerError()
	}
	payload := &operations.FiredUserOKBody{ID: row.ID(), Fired: *params.Body.Fired}
	log.OK(strconv.FormatUint(row.ID(), 10))
	return operations.NewFiredUserOK().WithPayload(payload)
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
	case params.Data.Email == nil:
		log.BadRequest("data.nick is null")
		return operations.NewCreateUserBadRequest()
	case params.Data.Surname == nil:
		log.BadRequest("data.surname is null")
		return operations.NewCreateUserBadRequest()
	case params.Data.Patronymic == nil:
		log.BadRequest("data.patronymic is null")
		return operations.NewCreateUserBadRequest()
	case params.Data.Tg == nil:
		log.BadRequest("data.tg is null")
		return operations.NewCreateUserBadRequest()
	case params.Data.Vk == nil:
		log.BadRequest("data.vk is null")
		return operations.NewCreateUserBadRequest()
	}
	//id, fail := uuid.Parse(params.Session)
	//if fail != nil {
	//	log.BadRequest("parse session id: %v", params.Session)
	//	return operations.NewCreateUserBadRequest()
	//}
	//session := c.sessions.Get(id)
	//if session != nil {
	//	log.BadRequest("session not found: %v", id)
	//	return operations.NewCreateUserBadRequest()
	//}
	//fmt.Println("createUser from", session.User().Name())
	apply, err := time.Parse(time.DateTime, params.Data.Apply)
	if err != nil {
		log.BadRequest("Invalid time format")
		return operations.NewCreateCallBadRequest()
	}
	row, fail := c.users.New(*params.Data.Name, *params.Data.Surname, *params.Data.Patronymic, *params.Data.Email,
		*params.Data.Vk, *params.Data.Tg, *params.Data.Nick, *params.Data.Password, *params.Data.Phone, &apply)
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

func (g GetUser) Handle(params operations.GetUserParams) middleware.Responder {
	log := g.log.Func("GetUser")
	row, fail := g.users.ByID(params.ID)
	if fail != nil {
		log.InternalSerer(fail.Error())
		return operations.NewListUsersInternalServerError()
	}
	payload := &operations.GetUserOKBody{ID: row.ID(), Name: row.Name(), Surname: row.Surname(),
		Patronymic: row.Patronymic(), Email: row.Email(), Vk: row.Vk(), Tg: row.Tg(), Nick: row.Nick(), Phone: row.Phone()}
	return operations.NewGetUserOK().WithPayload(payload)
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
		payload[index] = &operations.ListUsersOKBodyItems0{ID: item.ID(), Name: item.Name(), Surname: item.Surname(),
			Patronymic: item.Patronymic(), Email: item.Email(), Vk: item.Vk(), Tg: item.Tg(), Nick: item.Nick(),
			Phone: item.Phone()}
		if item.Apply() != nil {
			payload[index].Apply = item.Apply().String()
		}
		if item.Fired() != nil {
			payload[index].Fired = item.Fired().String()
		}
	}
	log.OK(strconv.Itoa(len(list)))
	return operations.NewListUsersOK().WithPayload(payload)
}

func (u UpdateUser) Handle(params operations.UpdateUserParams) middleware.Responder {
	log := u.log.Func("updateUser")
	switch {
	case params.Body.Name == nil:
		log.BadRequest("data.name is null")
		return operations.NewCreateUserBadRequest()
	case params.Body.Email == nil:
		log.BadRequest("data.nick is null")
		return operations.NewCreateUserBadRequest()
	case params.Body.Surname == nil:
		log.BadRequest("data.surname is null")
		return operations.NewCreateUserBadRequest()
	case params.Body.Patronymic == nil:
		log.BadRequest("data.patronymic is null")
		return operations.NewCreateUserBadRequest()
	case params.Body.Tg == nil:
		log.BadRequest("data.tg is null")
		return operations.NewCreateUserBadRequest()
	case params.Body.Vk == nil:
		log.BadRequest("data.vk is null")
		return operations.NewCreateUserBadRequest()
	}
	apply, err := time.Parse(time.DateTime, *params.Body.Apply)
	if err != nil {
		log.BadRequest("Invalid time format")
		return operations.NewCreateCallBadRequest()
	}
	row, fail := u.users.Update(params.ID, *params.Body.Name, *params.Body.Surname, *params.Body.Patronymic, *params.Body.Email,
		*params.Body.Vk, *params.Body.Tg, *params.Body.Nick, *params.Body.Phone, &apply)
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
