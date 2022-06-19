package models

import (
	"github.com/ShauryaAg/tcms-go/pkg/models"
)

func NewTestStep[T any](id T, description string) *TestStep[T] {
	model := models.NewModel(id)
	return &TestStep[T]{
		Model: *model,

		Description: description,
	}
}

type TestStep[T any] struct {
	models.Model[T] `bson:"inline"`

	Description string         `json:"description" bson:"description"`
	Parent      *TestStep[T]   `json:"parent" bson:"parent"`
	Children    []*TestStep[T] `json:"children" bson:"children"` // A step can have multiple child steps, depicting a tree
}
