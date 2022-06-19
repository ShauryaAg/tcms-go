package models

import (
	"github.com/ShauryaAg/tcms-go/pkg/models"
)

func NewTestCase[T any](id T, name string, priority string, testType string, version int, lastEditedBy string) *TestCase[T] {
	model := models.NewModel(id)
	return &TestCase[T]{
		Model: *model,

		Name:         name,
		Priority:     priority,
		Type:         testType,
		Version:      version,
		LastEditedBy: lastEditedBy,
	}
}

type TestCase[T any] struct {
	models.Model[T] `bson:"inline"`

	Name      string       `json:"name" bson:"name"`
	Priority  string       `json:"priority" bson:"priority"`
	Type      string       `json:"type" bson:"type"`
	FirstStep *TestStep[T] `json:"steps" bson:"steps"` // Stores the pointer to the first step

	Version      int    `json:"version" bson:"version"`
	LastEditedBy string `json:"last_edited_by" bson:"last_edited_by"`
}
