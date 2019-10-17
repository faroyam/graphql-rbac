package models

import (
	"errors"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

var (
	ErrIncorrectIDType = errors.New("incorrect ID type")
	ErrIncorrectID     = errors.New("incorrect ID")
)

type ID struct {
	UUID uuid.UUID
}

func IDFromUUID(id uuid.UUID) ID {
	return ID{UUID: id}
}

func IDFromUUIDPtr(id uuid.UUID) *ID {
	return &ID{UUID: id}
}

func (i *ID) UnmarshalGQL(v interface{}) error {

	idStr, ok := v.(string)
	if !ok {
		return ErrIncorrectIDType
	}
	id, err := uuid.FromString(idStr)
	if err != nil {
		return ErrIncorrectID
	}
	i.UUID = id
	return nil
}

func (i ID) MarshalGQL(w io.Writer) {
	graphql.MarshalID(i.UUID.String()).MarshalGQL(w)
	return
}
