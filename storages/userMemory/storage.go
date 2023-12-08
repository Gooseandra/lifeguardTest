package userMemory

import (
	"swagger/storages"
	"sync"
)

type (
	Storage struct {
		first *storageRow
		//list     []*storages.User
		dictById   storageRowsById
		dictByName storageRowsByName
		mutex      sync.Mutex
		sequence   storages.UserID
	}

	storageRow struct {
		next, prev *storageRow
		iD         storages.UserID
		name       storages.UserName
		password   storages.UserPassword
	}

	storageRowsById map[storages.UserID]*storageRow

	storageRowsByName map[storages.UserName]*storageRow
)

func NewStorage() *Storage {
	return &Storage{dictById: storageRowsById{}, dictByName: storageRowsByName{}}
}

func (s Storage) ByName(n storages.UserName) (storages.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if u, o := s.dictByName[n]; o {
		return u, nil
	}
	return nil, storages.UserNameMissingError(n)
}

func (s Storage) List(i uint64, c uint32) ([]storages.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	e, r := s.first, make([]storages.User, 0, c)
	for ; i > 0; i-- {
		if e == nil {
			return r, nil
		}
		e = e.next
	}
	for ; c > 0; c-- {
		if e == nil {
			return r, nil
		}
		r = append(r, e)
		e = e.next
	}
	return r, nil
}

func (s *Storage) New(n storages.UserName, p storages.UserPassword) (storages.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sequence++
	if r, o := s.dictById[s.sequence]; o {
		return r, storages.UserIdExistError(s.sequence)
	}
	if r, o := s.dictByName[n]; o {
		return r, storages.UserNameExistError(n)
	}
	r := &storageRow{iD: s.sequence, name: n, password: p}
	s.dictById[s.sequence], s.dictByName[n] = r, r
	l := &s.first
	for {
		if *l == nil {
			*l = r
			break
		} else if (*l).name < r.name {
			l = &(*l).next
		} else {
			r.next = *l
			*l = r
			break
		}
	}
	return r, nil
}

func (r *storageRow) ID() storages.UserID { return r.iD }

func (r *storageRow) Name() storages.UserName { return r.name }

func (r *storageRow) Password() storages.UserPassword { return r.password }
