package userPostgres

import (
	"database/sql"
	"swagger/storages"
)

const selectSql = `select "id","name","password","phone" from "user" limit $1 offset $2`
const newSql = `insert into "user" ("name", "password","phone")values($1,$2,$3) returning "id"`

type (
	Storage struct {
		db *sql.DB
	}

	storageRow struct {
		id       storages.UserID
		name     storages.UserName
		password storages.UserPassword
		phone    storages.UserPhone
	}
)

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s Storage) ByName(name storages.UserName) (storages.User, error) {
	panic("Not Implement")
}

func (s Storage) New(name, password, phone string) (storages.User, error) {
	row := s.db.QueryRow(newSql, name, password, phone)
	result := storageRow{name: name, password: password, phone: phone}
	err := row.Scan(&result.id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Storage) List(skip uint64, count uint32) ([]storages.User, error) {
	rows, err := s.db.Query(selectSql, count, skip)
	if err != nil {
		return nil, err
	}
	result := make([]storages.User, 0, count)
	for rows.Next() {
		var row storageRow
		err = rows.Scan(&row.id, &row.name, &row.password, &row.phone)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (r storageRow) ID() storages.UserID { return r.id }

func (r storageRow) Name() storages.UserName { return r.name }

func (r storageRow) Password() storages.UserPassword { return r.password }

func (r storageRow) Phone() storages.UserPhone { return r.phone }
