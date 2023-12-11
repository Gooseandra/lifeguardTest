package services

import (
	"fmt"
	"swagger/storages"
)

type Crews struct {
	storage storages.Crews
}

func NewCrew(s storages.Crews) *Crews { return &Crews{storage: s} }

func (c Crews) New(timeStart storages.CrewTime, leader storages.CrewLeader,
	comment storages.CrewComment) (storages.Crew, error) {
	crew, err := c.storage.New(timeStart, leader, comment)
	if err == nil {
		return crew, nil
	}
	return nil, err
}

func (c Crews) List(skip uint64, count uint32) ([]storages.Crew, error) {
	crew, err := c.storage.List(skip, count)
	if err == nil {
		return crew, nil
	}
	return nil, fmt.Errorf("user serice: %w", err)
}
