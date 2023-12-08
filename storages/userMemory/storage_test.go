package userMemory

import (
	"github.com/stretchr/testify/require"
	"swagger/storages"
	"testing"
)

func TestStorageByName(t *testing.T) {
	const n0, n1, n2, p1, p2 = "Zero", "One", "Two", "one", "two"
	notFound := func(t *testing.T) {
		var e storages.UserNameMissingError
		r1 := &storageRow{iD: 1, name: n1, password: p1}
		r2 := &storageRow{iD: 2, name: n2, password: p2}
		_, f := Storage{dictByName: storageRowsByName{n1: r1, n2: r2}}.ByName(n0)
		require.Error(t, f)
		require.ErrorAs(t, f, &e)
		require.Equal(t, e, storages.UserNameMissingError(n0))
	}
	success := func(t *testing.T) {
		r1 := &storageRow{iD: 1, name: n1, password: p1}
		r2 := &storageRow{iD: 2, name: n2, password: p2}
		a := Storage{dictByName: storageRowsByName{n1: r1, n2: r2}}
		r, f := a.ByName(n1)
		require.NoError(t, f)
		require.Equal(t, r, r1)
	}
	t.Run("not found", notFound)
	t.Run("success", success)
}

func TestStorageList(t *testing.T) {
	two := storageRow{iD: 2, name: "Two", password: "two"}
	three := storageRow{iD: 3, name: "Three", next: &two, password: "three"}
	one := storageRow{iD: 1, name: "One", next: &three, password: "one"}
	four := storageRow{iD: 4, name: "Four", next: &one, password: "four"}
	a := &Storage{first: &four}
	r, e := a.List(0, 9)
	require.NoError(t, e)
	require.Len(t, r, 4)
	require.Equal(t, r[0], &four)
	require.Equal(t, r[1], &one)
	require.Equal(t, r[2], &three)
	require.Equal(t, r[3], &two)
}

func TestStorageNew(t *testing.T) {
	const n1, n2, n3, n4, p1, p2, p3, p4 = "One", "Two", "Three", "Four", "One", "Two", "Three", "Four"
	duplicateId := func(t *testing.T) {
		var a storages.UserIdExistError
		s := Storage{dictById: storageRowsById{1: nil}}
		_, r := s.New(n1, p1)
		require.ErrorAs(t, r, &a)
		require.Equal(t, a, storages.UserIdExistError(1))
	}
	duplicateName := func(t *testing.T) {
		var e storages.UserNameExistError
		a := Storage{dictByName: storageRowsByName{n2: nil}}
		_, f := a.New(n2, p2)
		require.ErrorAs(t, f, &e)
		require.Equal(t, e, storages.UserNameExistError(n2))
	}
	success := func(t *testing.T) {
		type New struct {
			entity   storages.User
			name     storages.UserName
			password storages.UserPassword
		}
		a := Storage{dictById: storageRowsById{}, dictByName: storageRowsByName{}}
		n := []New{
			{name: n1, password: "one"},
			{name: n2, password: "two"},
			{name: n3, password: "three"},
			{name: n4, password: "four"}}
		for i, v := range n {
			var err error
			n[i].entity, err = a.New(v.name, v.password)
			require.NoError(t, err)
			require.Equal(t, storages.UserID(i+1), n[i].entity.ID())
			require.Equal(t, v.name, n[i].entity.Name())
			require.Equal(t, v.password, n[i].entity.Password())
		}
		require.Equal(t, storages.User(a.first), n[3].entity)
		require.NotNil(t, a.first.next)
		require.Equal(t, storages.User(a.first.next), n[0].entity)
		require.NotNil(t, a.first.next.next)
		require.Equal(t, storages.User(a.first.next.next), n[2].entity)
		require.NotNil(t, a.first.next.next.next)
		require.Equal(t, storages.User(a.first.next.next.next), n[1].entity)
		require.Nil(t, a.first.next.next.next.next)
		require.Equal(t, storages.User(a.dictById[1]), n[0].entity)
		require.Equal(t, storages.User(a.dictById[2]), n[1].entity)
		require.Equal(t, storages.User(a.dictById[3]), n[2].entity)
		require.Equal(t, storages.User(a.dictById[4]), n[3].entity)
		require.Equal(t, storages.User(a.dictByName[n1]), n[0].entity)
		require.Equal(t, storages.User(a.dictByName[n2]), n[1].entity)
		require.Equal(t, storages.User(a.dictByName[n3]), n[2].entity)
		require.Equal(t, storages.User(a.dictByName[n4]), n[3].entity)
	}
	t.Run("duplicateId", duplicateId)
	t.Run("duplicateName", duplicateName)
	t.Run("success", success)
}
