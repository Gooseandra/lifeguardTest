package storages

import "time"

type (
	Crew interface {
		ID() CrewID
		Start() CrewTime
		Finish() CrewTime
		Leader() CrewLeader
		Comment() CrewComment
		Roaster() CrewRoster
	}

	CrewID = uint64

	CrewTime = time.Time

	CrewLeader = uint64

	CrewComment = string

	CrewRoster = []uint64

	Crews interface {
		ByTime(time CrewTime) (Crew, error)
		New(start CrewTime, leader CrewLeader, comment CrewComment, roster CrewRoster) (Crew, error)
		List(skip uint64, count uint32) ([]Crew, error)
	}
)
