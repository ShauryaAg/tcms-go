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

func NewMongoStore(database string) *MongoStore {
	return &MongoStore{
		DatabaseName: database,
	}
}

type MongoStore struct {
	DatabaseName string
	database     *mongo.Database
	client       *mongo.Client
}

func (m *MongoStore) Connect(ctx context.Context) (interface{}, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionUri))
	if err != nil {
		return nil, err
	}

	m.database = client.Database(m.DatabaseName)
	m.client = client
	return client, nil
}

func (m *MongoStore) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *MongoStore) Database(ctx context.Context) interface{} {
	return m.database
}
