package mongosh

import (
	"context"
	"fmt"
	"product-serivce/internal/pkg/load"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoshStruct struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func ConnectMongo(cfg load.Config) (*MongoshStruct, error) {
	ctx := context.Background()
	target := fmt.Sprintf("mongodb://%s:%d", cfg.Mongosh.Host, cfg.Mongosh.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(target))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	collection := client.Database(cfg.Mongosh.Database).Collection(cfg.Mongosh.Collection)

	return &MongoshStruct{
		Client:     client,
		Collection: collection,
	}, nil
}
