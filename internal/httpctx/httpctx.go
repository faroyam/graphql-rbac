package httpctx

import (
	"context"

	"github.com/gofrs/uuid"
)

type Key string

const (
	AccessTokenKey       = Key("access_token")
	UserIDKey            = Key("user_id")
	AllowedAllActionsKey = Key("allowed_all_actions")
	AllowedActionsKey    = Key("allowed_actions")
)

func GetUser(ctx context.Context) uuid.UUID {
	id, _ := ctx.Value(UserIDKey).(uuid.UUID)
	return id
}
