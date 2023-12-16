package inventoryPostgres

import (
	"database/sql"
	"swagger/storages"
)

const newItemSql = `insert into "inventory"("name", "type", "description", "uniqNum")values($1,$2,$3,$4) returning "id"`
const selectListSql = `select "id", "name", "type", "description", "uniqNum" from "inventory" limit $1 offset $2`

type (
	Storage struct {
		db *sql.DB
	}

	storageRow struct {
		id          storages.IID
		typeName    storages.ITypeName
		name        storages.IName
		description storages.IDescription
		uniqNum     storages.IUniqNum
	}
)

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s Storage) New(typeName storages.ITypeName, name storages.IName,
	description storages.IDescription, uniqNum storages.IUniqNum) (storages.Inventory, error) {
	row := s.db.QueryRow(newItemSql, name, typeName, description, uniqNum)
	result := storageRow{name: name, typeName: typeName, description: description, uniqNum: uniqNum}
	err := row.Scan(&result.id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Storage) List(count uint32, skip uint64) ([]storages.Inventory, error) {
	rows, err := s.db.Query(selectListSql, count, skip)
	if err != nil {
		return nil, err
	}
	result := make([]storages.Inventory, 0, count)
	for rows.Next() {
		var row storageRow
		err = rows.Scan(&row.id, &row.name, &row.typeName, &row.description, &row.uniqNum)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (s storageRow) TypeName() storages.ITypeName { return s.typeName }

func (s storageRow) ID() storages.IID { return s.id }

func (s storageRow) Name() storages.IName { return s.name }

func (s storageRow) InstanceDesc() storages.IDescription { return s.description }

func (s storageRow) UniqNum() storages.IUniqNum { return s.uniqNum }
