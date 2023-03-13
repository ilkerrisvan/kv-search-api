package repository

import (
	"context"
	"fmt"
	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"kv-search-api/internal/model"
	"kv-search-api/pkg/config"
	"sync"
)

type Repository struct {
	inMemoryDB        map[string]string
	mu                *sync.Mutex
	mongoClient       *mongo.Client
	recordsCollection *mongo.Collection
}

func NewRepository(mongoClient *mongo.Client) *Repository {
	return &Repository{
		inMemoryDB:        make(map[string]string),
		mu:                &sync.Mutex{},
		mongoClient:       mongoClient,
		recordsCollection: config.GetCollection(mongoClient, "records"),
	}
}

/*
get records then sum counts
*/
func (r Repository) GetRecords(req model.RecordsRequest) ([]bson.M, error) {
	var resData []bson.M

	query := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gt": req.StartDate,
					"$lt": req.EndDate,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": req.MinCount,
					"$lt": req.MaxCount,
				},
			},
		},
	}

	ctx := context.TODO()
	cursor, err := r.recordsCollection.Aggregate(ctx, query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}(cursor, ctx)

	err = cursor.All(ctx, &resData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resData, nil
}

/*
checks is key using or not
*/
func (r *Repository) IsKeyUsedBefore(k string) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	if r.inMemoryDB[k] != "" {
		return true
	}
	return false
}

/*
saves key - value pair to memeory
*/
func (r *Repository) SetPair(k string, v string) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	r.inMemoryDB[k] = v
	return false
}

/*
returns the value of the key
*/
func (r *Repository) GetValue(k string) string {
	defer r.mu.Unlock()
	r.mu.Lock()
	if !(r.inMemoryDB[k] == "") {
		return r.inMemoryDB[k]
	}
	return ""
}
