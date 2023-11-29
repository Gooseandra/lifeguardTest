package storages

import "strconv"

type (
	User interface {
		ID() UserID
		Name() UserName
		Password() UserPassword
	}

	UserById map[UserID]User

	UserExistError UserID

	UserID = uint64

	UserName = string

	UserPassword = string

	Users interface {
		New(n, p UserName) (User, error)
		List(uint64, uint32) ([]User, error)
	}
)

func (e UserExistError) Error() string {
	return "User already exist: id=" + strconv.FormatUint(UserID(e) /* !!! */, 10)
}
