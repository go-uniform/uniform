package domain

import (
	"github.com/go-diary/diary"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// A definition of the public functions for a mongo interface
type INoSql interface {
	CatchErrNoResults(handler func(p diary.IPage))

	/* Utility */
	Count(timeout time.Duration, database, collection string, query bson.D) int64
	Inc(timeout time.Duration, database, collection string, query bson.D, amounts map[string]float64) ([]interface{}, int64)

	/* Single */
	FindOne(timeout time.Duration, database, collection string, sort string, skip int64, query bson.D, model interface{})
	InsertOne(timeout time.Duration, database, collection string, document interface{}, model interface{})
	UpdateOne(timeout time.Duration, database, collection string, identifier, document interface{}, model interface{})
	DeleteOne(timeout time.Duration, database, collection string, document interface{}, model interface{})

	/* Bulk */
	FindMany(timeout time.Duration, database, collection string, sort string, skip, limit int64, query bson.D, model interface{})
	InsertMany(timeout time.Duration, database, collection string, documents ...interface{}) ([]interface{}, int64)
	UpdateMany(timeout time.Duration, database, collection string, query bson.D, document interface{}) ([]interface{}, int64)
	DeleteMany(timeout time.Duration, database, collection string, query bson.D) ([]interface{}, int64)
}
