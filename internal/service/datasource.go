package service

import (
	"context"
	"fmt"
	"time"

	"github.com/eviltomorrow/robber-datasource/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "robber"
	collectionName = "metadata"
)

func InsertMetadataMany(db *mongo.Client, metadata []*model.Metadata, timeout time.Duration) (int64, error) {
	if len(metadata) == 0 {
		return 0, nil
	}
	var collection = db.Database(databaseName).Collection(collectionName)
	var data = make([]interface{}, 0, len(metadata))
	for _, md := range metadata {
		data = append(data, bson.M{
			"source":           "sina",
			"code":             md.Code,
			"name":             md.Name,
			"open":             md.Open,
			"yesterday_closed": md.YesterdayClosed,
			"high":             md.High,
			"low":              md.Low,
			"latest":           md.Latest,
			"volume":           md.Volume,
			"account":          md.Account,
			"date":             md.Date,
			"time":             md.Time,
			"suspend":          md.Suspend,
			"create_timestamp": time.Now().Unix(),
			"modify_timestamp": 0,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := collection.InsertMany(ctx, data)
	if err != nil {
		return 0, err
	}
	if result == nil {
		return 0, fmt.Errorf("panic: InsertMany result is nil")
	}
	return int64(len(result.InsertedIDs)), nil
}

func DeleteMetadataByDate(db *mongo.Client, code, date string, timeout time.Duration) (int64, error) {
	if date == "" {
		return 0, nil
	}
	var collection = db.Database(databaseName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := collection.DeleteMany(ctx, bson.M{
		"code": code,
		"date": date,
	})
	if err != nil {
		return 0, err
	}
	if result == nil {
		return 0, fmt.Errorf("panic: DeleteMany result is nil")
	}
	return result.DeletedCount, nil
}

func SelectMetadataRange(db *mongo.Client, offset, limit int64, date string, lastID string, timeout time.Duration) ([]*model.Metadata, error) {
	if date == "" {
		return nil, fmt.Errorf("invalid date")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("invalid limit")
	}

	ctx, cannel := context.WithTimeout(context.Background(), timeout)
	defer cannel()

	var collection = db.Database(databaseName).Collection(collectionName)
	var opt = &options.FindOptions{}
	opt.SetLimit(limit)

	var filter = bson.M{
		"date": date,
	}
	if lastID != "" {
		objectID, err := primitive.ObjectIDFromHex(lastID)
		if err != nil {
			return nil, err
		}
		filter = bson.M{"_id": bson.M{"$gt": objectID}, "date": date}
	} else {
		opt.SetSkip(offset)
	}

	cur, err := collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var data = make([]*model.Metadata, 0, limit)
	for cur.Next(context.Background()) {
		var m = &model.Metadata{}
		if err := cur.Decode(m); err != nil {
			return nil, err
		}
		data = append(data, m)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
