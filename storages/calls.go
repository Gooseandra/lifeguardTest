package storages

import "time"

type (
	Call interface {
		ID() CallID
		Description() CallDescription
		SummingUp() CallSummingUp
		Address() CallAddress
		Time() CallTime
	}

	CallID = uint64

	CallDescription = string

	CallSummingUp = string

	CallAddress = string

	CallTime = time.Time

	Calls interface {
		ByID(id CallID) (Call, error)
		New(description CallDescription, address CallAddress, time CallTime) (Call, error)
		List(skip uint64, count uint32) ([]Call, error)
	}
)
