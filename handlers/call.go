package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"strconv"
	"swagger/restapi/operations"
	"swagger/services"
	"time"
)

type (
	call struct {
		log      *services.Log
		sessions *services.Sessions
		calls    *services.Calls
	}
	CreateCall struct{ call }
	GetCall    struct{ call }
	ListCall   struct{ call }
	UpdateCall struct{ call }
)

func NewCreateCall(l *services.Log, s *services.Sessions, c *services.Calls) CreateCall {
	return CreateCall{call: call{log: l, sessions: s, calls: c}}
}

func NewListCall(l *services.Log, s *services.Sessions, c *services.Calls) ListCall {
	return ListCall{call: call{log: l, sessions: s, calls: c}}
}

func (c CreateCall) Handle(params operations.CreateCallParams) middleware.Responder {
	log := c.log.Func("createCrew")
	switch {
	case params.Body.TimeStart == nil:
		log.BadRequest("Time start is null")
		return operations.NewCreateCallBadRequest()
	case params.Body.Description == nil:
		log.BadRequest("description is null")
		return operations.NewCreateCallBadRequest()
	}
	timeStart, err := time.Parse(time.DateTime, *params.Body.TimeStart)
	if err != nil {
		log.BadRequest("Invalid time format")
		return operations.NewCreateCallBadRequest()
	}
	row, fail := c.calls.New(timeStart, *params.Body.Description, *params.Body.Address)
	if fail != nil {
		log.InternalSerer(fail.Error())
		return operations.NewCreateUserInternalServerError()
	}
	log.OK(strconv.FormatUint(row.ID(), 10))
	return operations.NewCreateUserOK().WithPayload(row.ID())
}
