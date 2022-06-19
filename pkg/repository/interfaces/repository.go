package interfaces

import "context"

type Repository interface {
	FindById(ctx context.Context, objectId interface{}) (interface{}, error)
	Find(ctx context.Context, filter interface{}) (interface{}, error)
	FindOne(ctx context.Context, filter interface{}) (interface{}, error)
	Create(ctx context.Context, data interface{}) (interface{}, error)
	Update(ctx context.Context, filter interface{}, model interface{}) error
	Delete(ctx context.Context, filter interface{}) error
}
