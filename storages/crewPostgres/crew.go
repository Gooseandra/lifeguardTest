package crewPostgres

import (
	"database/sql"
	"swagger/storages"
)

const crewNewSql = `insert into "day_crew"("time_start", "time_end", "leader", "comment")values ($1,$1,$2,$3) returning "id"`
const selectSql = `select "id","time_start","time_end","leader", "comment" from "day_crew" limit $1 offset $2`

type (
	Storage struct {
		db *sql.DB
	}

	storageRow struct {
		id         storages.CrewID
		timeStart  storages.CrewTime
		timeFinish storages.CrewTime
		leader     storages.CrewLeader
		comment    storages.CrewComment
		roster     storages.CrewRoster
	}
)

func (s Storage) ByTime(time storages.CrewTime) (storages.Crew, error) {
	panic("Not Implement")
}

func (s Storage) New(start storages.CrewTime, leader storages.CrewLeader,
	comment storages.CrewComment, roster storages.CrewRoster) (storages.Crew, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow(crewNewSql, start, leader, comment)
	result := storageRow{timeStart: start, leader: leader, comment: comment, roster: roster}
	err = row.Scan(&result.id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, item := range roster {
		_, err = tx.Exec(`insert into "day_crew_roster"("user", "day_crew")values($1,$2)`, item, result.id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return result, nil
}

func (s Storage) List(skip uint64, count uint32) ([]storages.Crew, error) {
	rows, err := s.db.Query(selectSql, count, skip)
	if err != nil {
		return nil, err
	}
	result := make([]storages.Crew, 0, count)
	for rows.Next() {
		var row storageRow
		err = rows.Scan(&row.id, &row.timeStart, &row.timeFinish, &row.leader, &row.comment)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (r storageRow) ID() storages.CrewID { return r.id }

func (r storageRow) Start() storages.CrewTime { return r.timeStart }

func (r storageRow) Finish() storages.CrewTime { return r.timeFinish }

func (r storageRow) Leader() storages.CrewLeader { return r.leader }

func (r storageRow) Comment() storages.CrewComment { return r.comment }

func (r storageRow) Roaster() storages.CrewRoster { return r.roster }
