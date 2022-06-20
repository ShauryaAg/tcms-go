package mongo

import (
	"context"
	"errors"

	models "github.com/ShauryaAg/tcms-go/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoTestCaseRepository(db *mongo.Database) *MongoTestCaseRepository {
	return &MongoTestCaseRepository{
		DB:         db,
		Collection: db.Collection("test_cases"),
	}
}

type MongoTestCase = models.TestCase[primitive.ObjectID]

type MongoTestCaseRepository struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

func (m MongoTestCaseRepository) FindById(ctx context.Context, objectId interface{}) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(objectId.(string))
	if err != nil {
		return MongoTestCase{}, err
	}

	return FindById[MongoTestCase](ctx, m.Collection, id)
}

func (m MongoTestCaseRepository) Find(ctx context.Context, filter interface{}) (interface{}, error) {
	return Find[MongoTestCase](ctx, m.Collection, filter)
}

func (m MongoTestCaseRepository) FindOne(ctx context.Context, filter interface{}) (interface{}, error) {
	return FindOne[MongoTestCase](ctx, m.Collection, filter)
}

func (m MongoTestCaseRepository) Create(ctx context.Context, data interface{}) (interface{}, error) {
	ma, ok := data.(map[string]interface{})
	if !ok {
		return MongoTestCase{}, errors.New("data must be a map[string]interface{}")
	}

	model := models.NewTestCase(
		primitive.NewObjectID(),
		ma["name"].(string),
		ma["priority"].(string),
		ma["testType"].(string),
		int(ma["version"].(float64)),
		ma["editedBy"].(string),
	)

	return Create[MongoTestCase](ctx, m.Collection, *model)
}

func (m MongoTestCaseRepository) Update(ctx context.Context, filter interface{}, data interface{}) error {
	return Update[MongoTestCase](ctx, m.Collection, filter, data)
}

func (m MongoTestCaseRepository) Delete(ctx context.Context, filter interface{}) error {
	return Delete[MongoTestCase](ctx, m.Collection, filter)
}
