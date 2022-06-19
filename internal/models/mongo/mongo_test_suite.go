package mongo

import (
	"github.com/ShauryaAg/tcms-go/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTestSuite struct {
	models.TestSuite[primitive.ObjectID] `bson:"inline"`

	collection *mongo.Collection
}

func (m MongoTestSuite) Collection() *mongo.Collection {
	return m.collection
}

func (m *MongoTestSuite) ToTestSuite() models.TestSuite[primitive.ObjectID] {
	return m.TestSuite
}
