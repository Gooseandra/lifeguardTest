package userMemory

import (
	"swagger/storages"
	"sync"
)

type (
	storage struct {
		first *storageRow
		//list     []*storages.User
		dict     storageRowsById
		mutex    sync.Mutex
		sequence storages.UserID
	}

	storageRow struct {
		next, prev *storageRow
		iD         storages.UserID
		name       storages.UserName
		password   storages.UserPassword
	}

	storageRowsById map[storages.UserID]*storageRow
)

func NewStorage() *storage { return &storage{dict: storageRowsById{}} }

func (s storage) List(i uint64, c uint32) ([]storages.User, error) {
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

func (s *storage) New(n storages.UserName, p storages.UserPassword) (storages.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sequence++
	if r, o := s.dict[s.sequence]; o {
		return r, storages.UserExistError(s.sequence)
	}
	r := &storageRow{iD: s.sequence, name: n, password: p}
	s.dict[s.sequence] = r
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

func (r storageRow) ID() storages.UserID { return r.iD }

func (r storageRow) Name() storages.UserName { return r.name }

func (r storageRow) Password() storages.UserPassword { return r.password }
