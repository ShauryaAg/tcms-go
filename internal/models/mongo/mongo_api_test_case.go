package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ShauryaAg/tcms-go/internal/models"
)

func NewMongoAPITestCase(name string, priority string, testType string, version int, editedBy string,
	method string, url string, endpoint string,
	inputHeaders interface{}, inputCookies interface{}, inputBody interface{},
	expectedStatus int, expectedHeaders interface{}, expectedCookies interface{}, expectedBody interface{},
) *MongoAPITestCase {
	id := primitive.NewObjectID()
	apiTestCase := models.NewAPITestCase(id, name, priority, version, editedBy, method, url, endpoint, inputHeaders, inputCookies, inputHeaders, expectedStatus, expectedBody, expectedHeaders, expectedCookies)

	return &MongoAPITestCase{
		APITestCase: *apiTestCase,
	}
}

type MongoAPITestCase struct {
	models.APITestCase[primitive.ObjectID] `bson:"inline"`

	collection *mongo.Collection
}

func (m MongoAPITestCase) Collection() *mongo.Collection {
	return m.collection
}

func (m *MongoAPITestCase) ToAPITestCase() models.APITestCase[primitive.ObjectID] {
	return m.APITestCase
}
