package models

import "time"

func NewModel[T any](id T) *Model[T] {
	return &Model[T]{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Model[T any] struct {
	ID        T         `json:"_id" bson:"_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
