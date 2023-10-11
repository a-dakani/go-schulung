package mongodb

import (
	"context"
	"github.com/a-dakani/go-schulung/http-server-gin-persistence-mongo/ginserver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AutoRepository struct {
	client *mongo.Client
}

func NewAutoRepository(ctx context.Context) (*AutoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &AutoRepository{client: client}, nil
}
func (ar *AutoRepository) Close() error {
	return ar.client.Disconnect(context.Background())
}

func (ar *AutoRepository) AddAuto(ctx context.Context, auto ginserver.Auto) error {
	autoCollection := ar.client.Database("autos").Collection("auto")
	_, err := autoCollection.InsertOne(ctx, auto)
	return err
}

func (ar *AutoRepository) GetAllAutos(ctx context.Context) ([]ginserver.Auto, error) {
	autoCollection := ar.client.Database("autos").Collection("auto")
	cursor, err := autoCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var autos []ginserver.Auto
	for cursor.Next(ctx) {
		var audi ginserver.Audi
		err := cursor.Decode(&audi)
		if err != nil {
			return nil, err
		}
		autos = append(autos, audi)
	}
	return autos, nil
}
