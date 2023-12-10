package storages

import "time"

type (
	Crew interface {
		ID() CrewID
		Start() CrewTime
		Finish() CrewTime
		Leader() CrewLeader
		Comment() CrewComment
	}

	CrewID = uint64

	CrewTime = time.Time

	CrewLeader = uint64

	CrewComment = string

	Crews interface {
		ByTime(time CrewTime) (Crew, error)
		New(start CrewTime, leader CrewLeader, comment CrewComment) (Crew, error)
		List(skip uint64, count uint32) ([]Crew, error)
	}
)
