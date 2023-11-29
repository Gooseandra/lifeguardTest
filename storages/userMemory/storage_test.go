package userMemory

import (
	"github.com/stretchr/testify/require"
	"swagger/storages"
	"testing"
)

func TestNewStorageList(t *testing.T) {
	two := storageRow{iD: 2, name: "Two", password: "two"}
	three := storageRow{iD: 3, name: "Three", next: &two, password: "three"}
	one := storageRow{iD: 1, name: "One", next: &three, password: "one"}
	four := storageRow{iD: 4, name: "Four", next: &one, password: "four"}
	a := storage{first: &four}
	r, e := a.List(0, 9)
	require.NoError(t, e)
	require.Len(t, r, 4)
	require.Equal(t, r[0], &four)
	require.Equal(t, r[1], &one)
	require.Equal(t, r[2], &three)
	require.Equal(t, r[3], &two)
}

func TestStorageNew(t *testing.T) {
	duplicate := func(t *testing.T) {
		var e = storages.UserExistError(1)
		a := storage{dict: storageRowsById{1: nil}}
		_, f := a.New("name", "password")
		require.ErrorAs(t, f, &e)
	}
	success := func(t *testing.T) {
		type New struct {
			entity   storages.User
			name     storages.UserName
			password storages.UserPassword
		}
		a := storage{dict: storageRowsById{}}
		n := []New{
			{name: "One", password: "one"},
			{name: "Two", password: "two"},
			{name: "Three", password: "three"},
			{name: "Four", password: "four"}}
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
		require.Equal(t, storages.User(a.dict[1]), n[0].entity)
		require.Equal(t, storages.User(a.dict[2]), n[1].entity)
		require.Equal(t, storages.User(a.dict[3]), n[2].entity)
		require.Equal(t, storages.User(a.dict[4]), n[3].entity)
	}
	t.Run("duplicate", duplicate)
	t.Run("success", success)
}
