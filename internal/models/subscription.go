package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Subscription struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	UserID      int64         `bson:"user_id"`
	Latitude    float64       `bson:"latitude"`
	Longitude   float64       `bson:"longitude"`
	TimeZone    string        `bson:"time_zone"`
	LocalTime   time.Time     `bson:"local_time"`
	NotifyAtUTC time.Time     `bson:"utc_time"`
}

type TempSubscription struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	UserID      int64         `bson:"user_id"`
	Latitude    float64       `bson:"latitude"`
	Longitude   float64       `bson:"longitude"`
	TimeZone    string        `bson:"time_zone"`
	LocalTime   time.Time     `bson:"local_time"`
	NotifyAtUTC time.Time     `bson:"utc_time"`
}
