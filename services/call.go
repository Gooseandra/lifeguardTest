package services

import (
	"swagger/storages"
)

type Calls struct {
	storage storages.Calls
}

func NewCall(s storages.Calls) *Calls { return &Calls{storage: s} }

func (c Calls) New(timeStart storages.CallTime, desc storages.CallDescription,
	address storages.CallAddress) (storages.Call, error) {
	call, err := c.storage.New(desc, address, timeStart)
	if err == nil {
		return call, nil
	}
	return nil, err
}

func (c Calls) List(skip uint64, count uint32) ([]storages.Call, error) {
	panic("Not Implement")
}
