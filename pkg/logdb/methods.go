package logdb

import (
	"log"

	"github.com/influenzanet/logging-service/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dbService *LogDBService) SaveLogEvent(instanceID string, event types.LogEvent) (string, error) {
	ctx, cancel := dbService.getContext()
	defer cancel()

	res, err := dbService.collectionRefLogs(instanceID).InsertOne(ctx, event)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (dbService *LogDBService) FindLogEvents(
	instanceID string,
	query types.LogQuery,
	cbk func(instanceID string, logEvent types.LogEvent, args ...interface{}) error,
	args ...interface{},
) (err error) {
	ctx, cancel := dbService.getContext()
	defer cancel()

	filter := bson.M{}
	if query.Start > 0 && query.End > 0 {
		filter["$and"] = bson.A{
			bson.M{"time": bson.M{"$gt": query.Start}},
			bson.M{"time": bson.M{"$lt": query.End}},
		}
	} else if query.Start > 0 {
		filter["time"] = bson.M{"$gt": query.Start}
	} else if query.End > 0 {
		filter["time"] = bson.M{"$lt": query.End}
	}
	if len(query.UserID) > 0 {
		filter["userID"] = query.UserID
	}
	if len(query.Origin) > 0 {
		filter["origin"] = query.Origin
	}
	if len(query.EventName) > 0 {
		filter["eventName"] = query.EventName
	}
	if len(query.EventType) > 0 {
		filter["eventType"] = query.EventType
	}

	cur, err := dbService.collectionRefLogs(instanceID).Find(
		ctx,
		filter,
	)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result types.LogEvent
		err := cur.Decode(&result)
		if err != nil {
			return err
		}

		if err := cbk(instanceID, result, args...); err != nil {
			log.Printf("Unexpected Error: %v", err)
		}
	}
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}
