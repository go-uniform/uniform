package uniform

import (
	"fmt"
	"github.com/go-diary/diary"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type connMongo struct {
	c IConn
	p diary.IPage
}

func (m *connMongo) CatchNoDocumentsErr(handler func(p diary.IPage)) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("%v", r)
			if assertedErr, ok := r.(error); ok {
				err = assertedErr
			}
			if err.Error() != mongo.ErrNoDocuments.Error() {
				panic(err)
			}
		}
	}()

	if handler != nil {
		handler(m.p)
	}
}

func (m *connMongo) Aggregate(timeout time.Duration, database, collection string, stages []M, model interface{}) {
	if err := m.c.Request(m.p, "mongo.aggregate", timeout, Request{
		Model: M{
			"database":   database,
			"collection": collection,
			"stages":     stages,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m *connMongo) Count(timeout time.Duration, database, collection string, query M) int64 {
	var response int64
	if err := m.c.Request(m.p, "mongo.count", timeout, Request{
		Model: M{
			"database":   database,
			"collection": collection,
			"query":      query,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(&response)
	}); err != nil {
		panic(err)
	}
	return response
}

func (m *connMongo) GroupCount(timeout time.Duration, database, collection, groupField string, query M) int64 {
	var response []struct {
		Out int64 `json:"out"`
	}
	if err := m.c.Request(m.p, "mongo.aggregate", timeout, Request{
		Model: M{
			"database":   database,
			"collection": collection,
			"stages": []M{
				{
					"$match": query,
				},
				{
					"$group": M{
						"_id":   fmt.Sprintf("$%s", groupField),
						"count": M{"$sum": 1},
					},
				},
				{
					"$group": M{
						"_id": nil,
						"out": M{"$sum": 1},
					},
				},
			},
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(&response)
	}); err != nil {
		panic(err)
	}

	if len(response) > 1 {
		panic("was not expecting more than one record")
	} else if len(response) == 0 {
		return 0
	}

	return response[0].Out
}

func (m *connMongo) Avg(timeout time.Duration, database, collection, field string, query M) float64 {
	var response []struct {
		Out float64 `json:"out"`
	}

	if err := m.c.Request(m.p, "mongo.aggregate", timeout, Request{
		Model: M{
			"database":   database,
			"collection": collection,
			"stages": []M{
				{
					"$match": query,
				},
				{
					"$group": M{
						"_id": nil,
						"out": M{"$avg": fmt.Sprintf("$%s", field)},
					},
				},
			},
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(&response)
	}); err != nil {
		panic(err)
	}

	if len(response) > 1 {
		panic("was not expecting more than one record")
	} else if len(response) == 0 {
		return 0
	}

	return response[0].Out
}

func (m *connMongo) CountMonthly(timeout time.Duration, database, collection string, query M) map[string]float64 {
	var response []struct {
		Id struct {
			Year  int64 `json:"year"`
			Month int64 `json:"month"`
		} `json:"id"`
		Out int64 `json:"out"`
	}
	m.Aggregate(timeout, database, collection, []M{
		{
			"$match": query,
		},
		{
			"$group": M{
				"_id": M{
					"year":  M{"$year": "$created-at"},
					"month": M{"$month": "$created-at"},
				},
				"out": M{"$sum": 1},
			},
		},
	}, &response)

	if len(response) == 0 {
		return map[string]float64{}
	}

	valueMap := map[string]float64{}
	for _, item := range response {
		valueMap[fmt.Sprintf("%d-%d", item.Id.Year, item.Id.Month)] = float64(item.Out)
	}
	return valueMap
}

func (m *connMongo) GroupCountMonthly(timeout time.Duration, database, collection, groupField, dateField string, fromDate time.Time, query, out M) map[string]float64 {
	if fromDate.Day() > 1 {
		fromDate = time.Date(fromDate.Year(), fromDate.Month()+1, 1, 0, 0, 0, 0, fromDate.Location())
	} else {
		fromDate = time.Date(fromDate.Year(), fromDate.Month(), 1, 0, 0, 0, 0, fromDate.Location())
	}
	query[dateField] = M{"$gte": fromDate}

	var response []struct {
		Id struct {
			Year  int64 `json:"year"`
			Month int64 `json:"month"`
		} `json:"id"`
		Out float64 `json:"out"`
	}
	m.Aggregate(timeout, database, collection, []M{
		{
			"$match": query,
		},
		{
			"$group": M{
				"_id": M{
					"employee": fmt.Sprintf("$%s", groupField),
					"year":     M{"$year": fmt.Sprintf("$%s", dateField)},
					"month":    M{"$month": fmt.Sprintf("$%s", dateField)},
				},
				"count": M{"$sum": 1},
			},
		},
		{
			"$group": M{
				"_id": M{
					"year":  "$_id.year",
					"month": "$_id.month",
				},
				"out": out,
			},
		},
	}, &response)

	if len(response) == 0 {
		return map[string]float64{}
	}

	valueMap := map[string]float64{}
	for _, item := range response {
		valueMap[fmt.Sprintf("%d-%d", item.Id.Year, item.Id.Month)] = item.Out
	}
	return valueMap
}

func (m *connMongo) AverageMonthly(timeout time.Duration, database, collection, dateField, valueField string, fromDate time.Time, query M) map[string]float64 {
	if fromDate.Day() > 1 {
		fromDate = time.Date(fromDate.Year(), fromDate.Month()+1, 1, 0, 0, 0, 0, fromDate.Location())
	} else {
		fromDate = time.Date(fromDate.Year(), fromDate.Month(), 1, 0, 0, 0, 0, fromDate.Location())
	}
	query[dateField] = M{"$gte": fromDate}

	var response []struct {
		Id struct {
			Year  int64 `json:"year"`
			Month int64 `json:"month"`
		} `json:"id"`
		Out float64 `json:"out"`
	}
	m.Aggregate(timeout, database, collection, []M{
		{
			"$match": query,
		},
		{
			"$group": M{
				"_id": M{
					"year":  M{"$year": fmt.Sprintf("$%s", dateField)},
					"month": M{"$month": fmt.Sprintf("$%s", dateField)},
				},
				"out": M{"$avg": fmt.Sprintf("$%s", valueField)},
			},
		},
	}, &response)

	if len(response) == 0 {
		return map[string]float64{}
	}

	valueMap := map[string]float64{}
	for _, item := range response {
		valueMap[fmt.Sprintf("%d-%d", item.Id.Year, item.Id.Month)] = item.Out
	}
	return valueMap
}

func (m *connMongo) FindMany(timeout time.Duration, database, collection, sort string, skip, limit int64, query M, model interface{}, fieldTags map[string][]string) {
	if err := m.c.Request(m.p, "mongo.find.many", timeout, Request{
		Context: M{
			"field-tags": fieldTags,
		},
		Model: M{
			"database":   database,
			"collection": collection,
			"sort":       sort,
			"skip":       skip,
			"limit":      limit,
			"query":      query,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m *connMongo) FindOne(timeout time.Duration, database, collection string, sort string, skip int64, query M, model interface{}, fieldTags map[string][]string) {
	m.FindOneX(timeout, database, collection, sort, skip, query, model, fieldTags, false)
}

func (m *connMongo) FindOneX(timeout time.Duration, database, collection string, sort string, skip int64, query M, model interface{}, fieldTags map[string][]string, includeDeleted bool) {
	if err := m.c.Request(m.p, "mongo.find.one", timeout, Request{
		Context: M{
			"field-tags": fieldTags,
		},
		Model: M{
			"database":        database,
			"collection":      collection,
			"include-deleted": includeDeleted,
			"sort":            sort,
			"skip":            skip,
			"query":           query,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m*connMongo) Delete(timeout time.Duration, database, collection, id string, soft bool, model interface{}, fieldTags map[string][]string) {
	if err := m.c.Request(m.p, "mongo.delete", timeout, Request{
		Context: M{
			"field-tags": fieldTags,
		},
		Model: M{
			"database":   database,
			"collection": collection,
			"id":         id,
			"soft":       soft,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m *connMongo) Inc(timeout time.Duration, database, collection, id, field string, amount float64, model interface{}, fieldTags map[string][]string) {
	if err := m.c.Request(m.p, "mongo.inc", timeout, Request{
		Context: M{
			"field-tags": fieldTags,
		},
		Model: M{
			"database":   database,
			"collection": collection,
			"id":         id,
			"field":      field,
			"amount":     amount,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m *connMongo) Index(timeout time.Duration, database, collection, name string) {
	if err := m.c.Request(m.p, "mongo.index", timeout, Request{
		Model: M{
			"database":   database,
			"collection": collection,
			"name":       name,
		},
	}, nil); err != nil {
		panic(err)
	}
}

func (m *connMongo) Insert(timeout time.Duration, database, collection string, document interface{}, model interface{}, fieldTags map[string][]string) {
	if err := m.c.Request(m.p, "mongo.insert", timeout, Request{
		Context: M{
			"field-tags": fieldTags,
		},
		Model: M{
			"database":   database,
			"collection": collection,
			"document":   document,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m *connMongo) Read(timeout time.Duration, database, collection, id string, model interface{}, fieldTags map[string][]string) {
	if err := m.c.Request(m.p, "mongo.read", timeout, Request{
		Context: M{
			"field-tags": fieldTags,
		},
		Model: M{
			"database":   database,
			"collection": collection,
			"id":         id,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m *connMongo) Restore(timeout time.Duration, database, collection, id string, model interface{}, fieldTags map[string][]string) {
	if err := m.c.Request(m.p, "mongo.restore", timeout, Request{
		Context: M{
			"field-tags": fieldTags,
		},
		Model: M{
			"database":   database,
			"collection": collection,
			"id":         id,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m *connMongo) Update(timeout time.Duration, database, collection, id string, document interface{}, model interface{}, fieldTags map[string][]string) {
	m.UpdateX(timeout, database, collection, id, document, model, fieldTags, false)
}

func (m *connMongo) UpdateX(timeout time.Duration, database, collection, id string, document interface{}, model interface{}, fieldTags map[string][]string, includeDeleted bool) {
	if err := m.c.Request(m.p, "mongo.update", timeout, Request{
		Context: M{
			"field-tags": fieldTags,
		},
		Model: M{
			"database":        database,
			"collection":      collection,
			"include-deleted": includeDeleted,
			"id":              id,
			"document":        document,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}

func (m *connMongo) UpdateMany(timeout time.Duration, database, collection string, query M, partial interface{}) (matched, modified, upserted int64, upsertedId interface{}) {
	var response struct {
		MatchedCount  int64       `json:"matched"`     // The number of documents matched by the filter.
		ModifiedCount int64       `json:"modified"`    // The number of documents modified by the operation.
		UpsertedCount int64       `json:"upserted"`    // The number of documents upserted by the operation.
		UpsertedID    interface{} `json:"upserted-id"` // The _id field of the upserted document, or nil if no upsert was done.
	}

	if err := m.c.Request(m.p, "mongo.update.many", timeout, Request{
		Model: M{
			"database":   database,
			"collection": collection,
			"query":      query,
			"partial":    partial,
		},
	}, func(r IRequest, p diary.IPage) {
		r.Read(&response)
	}); err != nil {
		panic(err)
	}

	return response.MatchedCount, response.ModifiedCount, response.UpsertedCount, response.UpsertedID
}
