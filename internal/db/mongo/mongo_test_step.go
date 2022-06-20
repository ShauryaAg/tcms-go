package mongo

import (
	"context"
	"errors"

	models "github.com/ShauryaAg/tcms-go/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoTestStepRepository(db *mongo.Database) *MongoTestStepRepository {
	return &MongoTestStepRepository{
		DB:         db,
		Collection: db.Collection("test_steps"),
	}
}

type MongoTestStep = models.TestStep[primitive.ObjectID]

type MongoTestStepRepository struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

func (m MongoTestStepRepository) FindById(ctx context.Context, objectId interface{}) (interface{}, error) {
	id, ok := objectId.(primitive.ObjectID)
	if !ok {
		return MongoTestStep{}, errors.New("objectId must be a primitive.ObjectID")
	}

	return FindById[MongoTestStep](ctx, m.Collection, id)
}

func (m MongoTestStepRepository) Find(ctx context.Context, filter interface{}) (interface{}, error) {
	return Find[MongoTestStep](ctx, m.Collection, filter)
}

func (m MongoTestStepRepository) FindOne(ctx context.Context, filter interface{}) (interface{}, error) {
	return FindOne[MongoTestStep](ctx, m.Collection, filter)
}

func (m MongoTestStepRepository) Create(ctx context.Context, data interface{}) (interface{}, error) {
	ma, ok := data.(map[string]interface{})
	if !ok {
		return MongoTestStep{}, errors.New("data must be a map[string]interface{}")
	}

	model := models.NewTestStep(
		primitive.NewObjectID(),
		ma["description"].(string),
	)

	return Create[MongoTestStep](ctx, m.Collection, *model)
}

func (m MongoTestStepRepository) Update(ctx context.Context, filter interface{}, data interface{}) error {
	return Update[MongoTestStep](ctx, m.Collection, filter, data)
}

func (m MongoTestStepRepository) Delete(ctx context.Context, filter interface{}) error {
	return Delete[MongoTestStep](ctx, m.Collection, filter)
}
