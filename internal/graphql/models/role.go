package models

import "graphql-rbac/internal/repository"

func RoleFromRepo(in *repository.Role) (*Role, error) {
	user := &Role{
		ID:          IDFromUUID(in.ID),
		Title:       in.Title,
		Description: in.Description,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return user, nil
}

func RolesFromRepo(in []*repository.Role) ([]*Role, error) {
	roles := make([]*Role, 0, len(in))
	for _, u := range in {
		user, err := RoleFromRepo(u)
		if err != nil {
			return nil, err
		}
		roles = append(roles, user)
	}
	return roles, nil
}
