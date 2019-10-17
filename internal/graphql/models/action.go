package models

import "graphql-rbac/internal/repository"

func ActionFromRepo(in *repository.Action) (*GraphqlAction, error) {
	action := &GraphqlAction{
		Title:       in.Title,
		Description: in.Description,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return action, nil
}

func ActionsFromRepo(in []*repository.Action) ([]*GraphqlAction, error) {
	actions := make([]*GraphqlAction, 0, len(in))
	for _, u := range in {
		user, err := ActionFromRepo(u)
		if err != nil {
			return nil, err
		}
		actions = append(actions, user)
	}
	return actions, nil
}
