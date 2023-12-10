package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"strconv"
	"swagger/restapi/operations"
	"swagger/services"
	"time"
)

type (
	crew struct {
		log      *services.Log
		sessions *services.Sessions
		crew     *services.Crew
	}
	CreateCrew struct{ crew }
	GetCrew    struct{ crew }
	ListCrew   struct{ crew }
	UpdateCrew struct{ crew }
)

func NewCreateCrew(l *services.Log, s *services.Sessions, c *services.Crew) CreateCrew {
	return CreateCrew{crew: crew{log: l, sessions: s, crew: c}}
}

func NewGetCrew(l *services.Log, s *services.Sessions, c *services.Crew) GetCrew {
	return GetCrew{crew: crew{log: l, sessions: s, crew: c}}
}

func NewListCrew(l *services.Log, s *services.Sessions, c *services.Crew) ListCrew {
	return ListCrew{crew: crew{log: l, sessions: s, crew: c}}
}

func NewUpdateCrew(l *services.Log, s *services.Sessions, c *services.Crew) UpdateCrew {
	return UpdateCrew{crew: crew{log: l, sessions: s, crew: c}}
}

func (c CreateCrew) Handle(params operations.CreateCrewParams) middleware.Responder {
	log := c.log.Func("createCrew")
	switch {
	case params.Body.TimeStart == nil:
		log.BadRequest("Time start is null")
		return operations.NewCreateUserBadRequest()
	case params.Body.Leader == nil:
		log.BadRequest("leader is null")
		return operations.NewCreateUserBadRequest()
	}
	timeStart, err := time.Parse(time.DateTime, *params.Body.TimeStart)
	if err != nil {
		log.BadRequest("Invalid time format")
		return operations.NewCreateCrewBadRequest()
	}
	row, fail := c.crew.New(*params.Data.Name, *params.Data.Password, *params.Data.Phone)
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

func (g GetCrew) Handle(p operations.GetUserParams) middleware.Responder {
	panic("Not Implement")
}

func (l ListCrew) Handle(p operations.ListUsersParams) middleware.Responder {
	//log := l.log.Func("listUsers")
	panic("Not Implement")
}

func (g UpdateCrew) Handle(p operations.UpdateUserParams) middleware.Responder {
	panic("Not Implement")
}
