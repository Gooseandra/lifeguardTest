package services

import (
	"fmt"
	"swagger/storages"
)

type Users struct{ storage storages.Users }

func NewUser(s storages.Users) Users { return Users{storage: s} }

func (u *Users) List(s uint64, c uint32) ([]storages.User, error) {
	users, err := u.storage.List(s, c)
	if err == nil {
		return users, nil
	}
	return nil, fmt.Errorf("user serice: %w", err)
}

func (u *Users) New(n, p string) (storages.User, error) {
	user, err := u.storage.New(n, p)
	if err == nil {
		return user, nil
	}
	return nil, fmt.Errorf("user serice: %w", err)
}
