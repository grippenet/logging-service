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
