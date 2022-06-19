package db

import (
	"context"
)

type Store interface {
	Connect(ctx context.Context) (interface{}, error)
	Disconnect(ctx context.Context) error
	Database(ctx context.Context) interface{}
}
