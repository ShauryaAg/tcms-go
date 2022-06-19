package models

import "github.com/ShauryaAg/tcms-go/pkg/models"

func NewTestSuite[T any](id T, name string, testCases []*TestCase[T], path []string) *TestSuite[T] {
	model := models.NewModel(id)
	return &TestSuite[T]{
		Model: *model,

		Name:      name,
		TestCases: testCases,
		Path:      path,
	}
}

type TestSuite[T any] struct {
	models.Model[T] `bson:"inline"`

	Name      string         `json:"name" bson:"name"`
	TestCases []*TestCase[T] `json:"test_cases" bson:"test_cases"`
	Path      []string       `json:"path" bson:"path"`
}
