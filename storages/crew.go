package storages

import "time"

type (
	Crew interface {
		ID() CrewID
		Start() CrewTime
		Finish() CrewTime
		Leader() CrewLeader
		Comment() CrewComment
		Roaster() CrewRoaster
	}

	CrewID = uint64

	CrewTime = time.Time

	CrewLeader = uint64

	CrewComment = string

	CrewRoaster = []int64

	Crews interface {
		ByTime(time CrewTime) (Crew, error)
		New(start CrewTime, leader CrewLeader, comment CrewComment) (Crew, error)
		List(skip uint64, count uint32) ([]Crew, error)
	}
)
