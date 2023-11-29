package impl

import (
	"github.com/go-openapi/runtime/middleware"
	"swagger/restapi/operations"
	"swagger/services"
)

type (
	CreateUser struct{ service *services.Users }
	ListUser   struct{ service *services.Users }
)

func NewCreateUser(u *services.Users) CreateUser { return CreateUser{service: u} }

func NewListUser(u *services.Users) ListUser { return ListUser{service: u} }

func (c CreateUser) Handle(p operations.CreateUserParams) middleware.Responder {
	if p.Data == nil || p.Data.Name == nil || p.Data.Password == nil {
		return operations.NewCreateUserBadRequest()
	}
	r, e := c.service.New(*p.Data.Name, *p.Data.Password)
	switch {
	case e == nil:
		return operations.NewCreateUserOK().WithPayload(r.ID())
	}
	return operations.NewCreateUserInternalServerError()
}

func (l ListUser) Handle(p operations.ListUsersParams) middleware.Responder {
	if p.Count == nil || p.Start == nil {
		return operations.NewListUsersBadRequest()
	}
	r, e := l.service.List(*p.Start, *p.Count)
	switch {
	case e == nil:
		payload := make([]*operations.ListUsersOKBodyItems0, len(r))
		for i, v := range r {
			payload[i] = &operations.ListUsersOKBodyItems0{ID: v.ID(), Name: v.Name()}
		}
		return operations.NewListUsersOK().WithPayload(payload)
	}
	return operations.NewListUsersInternalServerError()
}
