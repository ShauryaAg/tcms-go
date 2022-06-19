package mongo

import (
	"context"

	"github.com/ShauryaAg/tcms-go/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	store := db.NewMongoStore("tcms")
	store.Connect(context.Background())
	database := store.Database(context.Background()).(*mongo.Database)

	db.DbStore = store
	db.Repositories["test_cases"] = NewMongoTestCaseRepository(database)
}
