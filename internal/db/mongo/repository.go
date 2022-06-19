package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindById[T any](ctx context.Context, collection *mongo.Collection, objectId primitive.ObjectID) (T, error) {
	var model T
	collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&model)

	return model, nil
}

func Find[T any](ctx context.Context, collection *mongo.Collection, filter interface{}) ([]T, error) {
	var models []T
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, models); err != nil {
		return nil, err
	}

	return models, nil
}

func FindOne[T any](ctx context.Context, collection *mongo.Collection, filter interface{}) (T, error) {
	var model T
	if err := collection.FindOne(ctx, filter).Decode(&model); err != nil {
		return model, err
	}

	return model, nil
}

func Create[T any](ctx context.Context, collection *mongo.Collection, model T) (T, error) {
	_, err := collection.InsertOne(ctx, model)
	if err != nil {
		var null T
		return null, err
	}

	return model, err
}

func Update[T any](ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) error {
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func Delete[T any](ctx context.Context, collection *mongo.Collection, filter interface{}) error {
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
