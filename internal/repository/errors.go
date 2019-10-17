package repository

import "errors"

var (
	ErrNotFound               = errors.New("not found")
	ErrUnknown                = errors.New("unknown error")
	ErrUserIDIsTaken          = errors.New("user id is taken")
	ErrUserLoginIsTaken       = errors.New("user login is taken")
	ErrRoleIDIsTaken          = errors.New("role id is taken")
	ErrRoleTitleIsTaken       = errors.New("role title is taken")
	ErrUserAlreadyBoundToRole = errors.New("user already bound to role")
	ErrActionTitleIsTaken     = errors.New("action title is taken")

	ErrUserDoesNotExist   = errors.New("user does not exist")
	ErrRoleDoesNotExist   = errors.New("role does not exist")
	ErrActionDoesNotExist = errors.New("action does not exist")
)

func RepoErrors() map[string]error {
	return map[string]error{
		"users_id_uindex":                    ErrUserIDIsTaken,
		"users_login_uindex":                 ErrUserLoginIsTaken,
		"roles_id_uindex":                    ErrRoleIDIsTaken,
		"roles_title_uindex":                 ErrRoleTitleIsTaken,
		"users_roles_user_id_role_id_uindex": ErrUserAlreadyBoundToRole,
		"actions_pk":                         ErrActionTitleIsTaken,
		"actions_title_uindex":               ErrActionTitleIsTaken,

		"users_roles_user_id_fk":     ErrUserDoesNotExist,
		"users_roles_role_id_fk:":    ErrRoleDoesNotExist,
		"actions_roles_action_id_fk": ErrActionDoesNotExist,
		"actions_roles_role_id_fk":   ErrRoleDoesNotExist,
	}
}
