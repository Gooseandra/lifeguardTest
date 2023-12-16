package services

import (
	"fmt"
	"swagger/storages"
)

type Inventories struct {
	storage storages.Inventories
}

func NewInventory(s storages.Inventories) *Inventories { return &Inventories{storage: s} }

func (i Inventories) Create(typeName storages.ITypeName, name storages.IName,
	description storages.IDescription, uniqNum storages.IUniqNum) (storages.Inventory, error) {
	inventory, err := i.storage.New(typeName, name, description, uniqNum)
	if err == nil {
		return inventory, nil
	}
	return nil, fmt.Errorf("inventory serice: %w", err)
}

func (i Inventories) List(c uint32, s uint64) ([]storages.Inventory, error) {
	inventory, err := i.storage.List(c, s)
	if err == nil {
		return inventory, nil
	}
	return nil, fmt.Errorf("inventory serice: %w", err)
}
