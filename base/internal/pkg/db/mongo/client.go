package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"{project}/internal/pkg/db"
	"strconv"
	"time"
)

var mongoClient MongoClient

type MongoClient struct {
	client   *mongo.Client
	database *mongo.Database
}

// Create new mongodb client
func NewClient(config db.Configuration) (MongoClient, error) {
	m := MongoClient{}

	var connectionString string
	if config.Username != "" && config.Password != "" {
		connectionString = "mongodb://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + strconv.Itoa(config.Port)
	} else {
		connectionString = "mongodb://" + config.Host + ":" + strconv.Itoa(config.Port)
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(connectionString)
	clientOptions.SetConnectTimeout(time.Duration(config.Timeout) * time.Millisecond)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Timeout)*time.Millisecond)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return m, err
	}

	// Check the connection
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return m, err
	}

	m.client = client
	m.database = client.Database(config.DatabaseName)

	mongoClient = m

	return m, nil
}

// Clone mongodb client
func (mc MongoClient) CloseSession() {
	if mc.client != nil {
		mc.client.Disconnect(context.TODO())
		mc.client = nil
	}
}
