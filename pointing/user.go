package pointing

import "github.com/google/uuid"

type User struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func newUser(name string) (*User, error) {
	return &User{
		uuid.New().String(),
		name,
	}, nil
}
