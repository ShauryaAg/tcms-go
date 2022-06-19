package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri = "mongodb://mongoadmin:secret@localhost:27888"
)

var (
	connectionUri = uri
)

type MongoStore struct {
	DB *mongo.Database
}

func (m *MongoStore) Connect(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionUri))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *MongoStore) Disconnect(ctx context.Context) error {
	return nil
}

func (m *MongoStore) Database(ctx context.Context) interface{} {
	return m.DB
}
