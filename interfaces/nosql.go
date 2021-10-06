package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// A definition of the public functions for a mongo interface
type INoSql interface {
	Count(timeout time.Duration, database, collection string, query bson.D) int64
	InsertOne(timeout time.Duration, database, collection string, document interface{}, model interface{})
	FindOne(timeout time.Duration, database, collection string, sort string, skip int64, query bson.D, model interface{})

	//CatchNoDocumentsErr(handler func(p diary.IPage))
	//Aggregate(timeout time.Duration, database, collection string, stages []M, model interface{})
	//GroupCount(timeout time.Duration, database, collection, groupField string, query M) int64
	//Avg(timeout time.Duration, database, collection, field string, query M) float64
	//CountMonthly(timeout time.Duration, database, collection string, query M) map[string]float64
	//GroupCountMonthly(timeout time.Duration, database, collection, groupField, dateField string, fromDate time.Time, query, out M) map[string]float64
	//AverageMonthly(timeout time.Duration, database, collection, dateField, valueField string, fromDate time.Time, query M) map[string]float64
	//FindMany(timeout time.Duration, database, collection, sort string, skip, limit int64, query M, model interface{}, fieldTags map[string][]string)
	//Delete(timeout time.Duration, database, collection, id string, soft bool, model interface{}, fieldTags map[string][]string)
	//DeleteMany(timeout time.Duration, database, collection string, query M, soft bool) (deleted int64)
	//Inc(timeout time.Duration, database, collection, id, field string, amount float64, model interface{}, fieldTags map[string][]string)
	//Index(timeout time.Duration, database, collection, name string)
	//Read(timeout time.Duration, database, collection, id string, model interface{}, fieldTags map[string][]string)
	//Restore(timeout time.Duration, database, collection, id string, model interface{}, fieldTags map[string][]string)
	//Update(timeout time.Duration, database, collection, id string, document interface{}, model interface{}, fieldTags map[string][]string)
	//UpdateX(timeout time.Duration, database, collection, id string, document interface{}, model interface{}, fieldTags map[string][]string, includeDeleted bool)
	//UpdateMany(timeout time.Duration, database, collection string, query M, partial interface{}) (matched, modified, upserted int64, upsertedId interface{})
}
