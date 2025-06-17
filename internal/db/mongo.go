package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"

	"foxminded/3.3-weather-forecast-bot/internal/models"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

//var Collection *mongo.Collection

type MongoDBI interface {
	CreateSubscription(ctx context.Context, sub models.Subscription) error
	Exists(ctx context.Context, sub models.Subscription) (bool, error)
	DeleteSubscriptions(ctx context.Context, userID int64) error
	DeleteSubscription(ctx context.Context, id string) error
	GetUserSubscriptions(ctx context.Context, userID int64) ([]models.Subscription, error)
	GetDueSubscriptions(ctx context.Context, t time.Time) ([]models.Subscription, error)
	Close(ctx context.Context) error
}
type MongoDB struct {
	Collection *mongo.Collection
	Client     *mongo.Client
}

func NewMongoDB(ctx context.Context, uri string) (*MongoDB, error) {
	slogger.Log.Info("Connecting to MongoDB...")
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}
	db := client.Database("weatherbot")
	return &MongoDB{Collection: db.Collection("subscription"), Client: client}, nil
}

func (m *MongoDB) CreateSubscription(ctx context.Context, sub models.Subscription) error {
	_, err := m.Collection.InsertOne(ctx, sub)
	if err != nil {
		return fmt.Errorf("failed to write to db: %w", err)
	}
	return nil
}

func (m *MongoDB) Exists(ctx context.Context, sub models.Subscription) (bool, error) {
	filter := bson.M{
		"user_id":   sub.UserID,
		"latitude":  sub.Latitude,
		"longitude": sub.Longitude,
		"utc_time":  sub.NotifyAtUTC,
	}

	count, err := m.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("err %w", err)
	}
	return count > 0, nil
}

func (m *MongoDB) DeleteSubscriptions(ctx context.Context, userID int64) error {
	_, err := m.Collection.DeleteMany(ctx, bson.M{"user_id": userID})
	if err != nil {
		return fmt.Errorf("failed to delete subscriptions: %w", err)
	}
	return nil
}

func (m *MongoDB) DeleteSubscription(ctx context.Context, id string) error {
	bsonID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to parse bson id: %w", err)
	}
	_, err = m.Collection.DeleteOne(ctx, bson.M{"_id": bsonID})
	if err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}
	return nil
}

func (m *MongoDB) GetUserSubscriptions(ctx context.Context, userID int64) ([]models.Subscription, error) {
	cursor, err := m.Collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("failed to find subscriptions by user_id %d: %w", userID, err)
	}
	defer cursor.Close(ctx)

	var subscriptions []models.Subscription
	for cursor.Next(ctx) {
		var sub models.Subscription
		raw := cursor.Current
		slogger.Log.Debug("Unmarshalling subscription", "data", raw)
		if err = cursor.Decode(&sub); err != nil {
			return nil, fmt.Errorf("failed to decode subscription: %w", err)
		}
		subscriptions = append(subscriptions, sub)
		slogger.Log.Debug("Found subscription", "subscription", sub)
	}
	return subscriptions, nil
}

func (m *MongoDB) GetDueSubscriptions(ctx context.Context, t time.Time) ([]models.Subscription, error) {
	filter := bson.M{
		"$expr": bson.M{
			"$and": []bson.M{
				{"$eq": []interface{}{bson.M{"$hour": "$utc_time"}, t.UTC().Hour()}},
				{"$eq": []interface{}{bson.M{"$minute": "$utc_time"}, t.UTC().Minute()}},
			},
		},
	}

	cursor, err := m.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var subs []models.Subscription
	if err := cursor.All(ctx, &subs); err != nil {
		return nil, err
	}
	return subs, nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
