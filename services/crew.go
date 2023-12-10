package services

import (
	"swagger/storages"
)

type Crews struct {
	storage storages.Crew
}

func NewCrew(s storages.Crew) *Crews { return &Crews{storage: s} }
