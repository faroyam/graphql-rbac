package models

import (
	"graphql-rbac/internal/repository"
)

func UserFromRepo(in *repository.User) (*User, error) {
	user := &User{
		ID:        IDFromUUID(in.ID),
		Login:     in.Login,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return user, nil
}

func UsersFromRepo(in []*repository.User) ([]*User, error) {
	users := make([]*User, 0, len(in))
	for _, u := range in {
		user, err := UserFromRepo(u)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
