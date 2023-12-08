package services

import (
	"github.com/google/uuid"
	"log"
	"swagger/storages"
	"sync"
	"time"
)

type (
	SessionEntity struct {
		next, prev *SessionEntity
		iD         uuid.UUID
		time       time.Time
		user       storages.User
	}

	SessionId = uuid.UUID

	Sessions struct {
		duration    time.Duration
		event       chan struct{}
		log         *Log
		mutex       sync.Mutex
		first, last *SessionEntity
		dict        map[uuid.UUID]*SessionEntity
		users       *Users
	}
)

func NewSessions(l *Log, s *Users, d time.Duration) Sessions {
	result := Sessions{dict: map[uuid.UUID]*SessionEntity{}, duration: d, log: l, users: s}
	go result.routine()
	return result
}

func (e SessionEntity) ID() SessionId { return e.iD }

func (e SessionEntity) User() storages.User { return e.user }

func (s *Sessions) erase() {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	log.Printf("Sessions.erase 1\n")
	for s.first != nil && time.Now().After(s.first.time.Add(s.duration)) {
		log.Printf("Sessions.erase 2: id=%v time=%v\n", s.first.iD, s.first.time)
		delete(s.dict, s.first.iD)
		s.first = s.first.next
		if s.first == nil {
			s.last = nil
		}
	}
}

func (s Sessions) Get(i SessionId) *SessionEntity {
	defer s.erase()
	defer s.mutex.Unlock()
	s.mutex.Lock()
	if r, o := s.dict[i]; o {
		r.time = time.Now()
		return r
	}
	return nil
}

func (s *Sessions) New(u storages.User) (*SessionEntity, error) {
	s.erase()
	for {
		s.mutex.Lock()
		// TODO: нужен счетчик
		i := uuid.New()
		if _, o := s.dict[i]; !o {
			entity := &SessionEntity{next: nil, prev: s.last, iD: i, time: time.Now(), user: u}
			log.Printf("Sessions.New id=%v time=%v\n", entity.iD, entity.time)
			s.dict[i] = entity
			if s.last == nil {
				s.first, s.last = entity, entity
			} else {
				s.last, s.last.next = entity, entity
			}
			s.mutex.Unlock()
			return entity, nil
		}
		s.mutex.Unlock()
	}
}

func (s *Sessions) routine() {
	for {
		first := s.first
		log.Println("Sessions.routine 1")
		if first == nil {
			timer := time.After(s.duration)
			<-timer
		} else {
			timer := time.After(time.Now().Sub(first.time))
			<-timer
			s.erase()
		}
	}
}

func (s Sessions) Users() *Users { return s.users }
