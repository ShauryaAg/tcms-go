package mongo

import (
	"github.com/ShauryaAg/tcms-go/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTestStep struct {
	models.TestStep[primitive.ObjectID] `bson:"inline"`

	collection *mongo.Collection
}

func (m MongoTestStep) Collection() *mongo.Collection {
	return m.collection
}

func (m *MongoTestStep) ToTestStep() models.TestStep[primitive.ObjectID] {
	return m.TestStep
}
