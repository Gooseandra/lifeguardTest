package callPostgres

import (
	"database/sql"
	"swagger/storages"
)

const newCall = `insert into "calls"("time_start", "address", "desc") values($1,$2,$3) returning "id"`

type (
	Storage struct {
		db *sql.DB
	}

	storageRow struct {
		id          storages.CallID
		description storages.CallDescription
		summingUp   storages.CallSummingUp
		address     storages.CallAddress
		timeStart   storages.CallTime
		timeFinish  storages.CallTime
	}
)

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

//func (s Storage) ByID(id storages.CallID) (storages.Call, error){
//	panic("Not Implement")
//}

func (s Storage) New(description storages.CallDescription, address storages.CallAddress,
	time storages.CallTime) (storages.Call, error) {
	row := s.db.QueryRow(newCall, time, address, description)
	result := storageRow{timeStart: time, address: address, description: description}
	err := row.Scan(&result.id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Storage) List(skip uint64, count uint32) ([]storages.Call, error) {
	//rows, err := s.db.Query(selectSql, count, skip)
	//if err != nil {
	//	return nil, err
	//}
	//result := make([]storages.Crew, 0, count)
	//for rows.Next() {
	//	var row storageRow
	//	err = rows.Scan(&row.id, &row.timeStart, &row.timeFinish, &row.leader, &row.comment)
	//	if err != nil {
	//		return nil, err
	//	}
	//	result = append(result, row)
	//}
	//return result, nil
	panic("Not Implement")
}

func (s Storage) ByID(id storages.CallID) (storages.Call, error) {
	panic("Not Implement")
}

func (r storageRow) ID() storages.CallID { return r.id }

func (r storageRow) Description() storages.CallDescription { return r.description }

func (r storageRow) SummingUp() storages.CallSummingUp { return r.summingUp }

func (r storageRow) Address() storages.CallAddress { return r.address }

func (r storageRow) Time() storages.CallTime { return r.timeStart }

func (r storageRow) TimeFinish() storages.CallTime { return r.timeFinish }
