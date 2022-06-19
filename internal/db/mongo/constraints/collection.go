package constraints

import "go.mongodb.org/mongo-driver/mongo"

type Collection interface {
	Collection() *mongo.Collection
}
