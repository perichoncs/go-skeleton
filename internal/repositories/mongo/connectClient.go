package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectClient(dbURI string) (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbURI))

	if err != nil {
		return nil, err
	}

	// Ping to the db to check if the connection has been established successfully
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, err
}
