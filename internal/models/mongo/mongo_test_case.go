package mongo

import (
	"github.com/ShauryaAg/tcms-go/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoTestCase(name string, priority string, testType string, version int, editedBy string) *MongoTestCase {
	id := primitive.NewObjectID()
	testCase := models.NewTestCase(id, name, priority, testType, version, editedBy)

	return &MongoTestCase{
		TestCase: *testCase,
	}
}

type MongoTestCase struct {
	models.TestCase[primitive.ObjectID] `bson:"inline"`

	collection *mongo.Collection
}

func (m MongoTestCase) Collection() *mongo.Collection {
	return m.collection
}

func (m MongoTestCase) ToTestCase() models.TestCase[primitive.ObjectID] {
	return m.TestCase
}
