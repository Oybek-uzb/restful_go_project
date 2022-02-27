package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"restful_go_project/internal/user"
	"restful_go_project/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("creating a user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		d.logger.Errorf("failed to create user due to error: %v", err)
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}
	d.logger.Debug("converting insertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		d.logger.Errorf("failed to convert insertedID to ObjectID")
		return "", fmt.Errorf("failed to convert ObjectID")
	}
	d.logger.Info("user created")
	return oid.Hex(), nil
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert ObjectID to Hex error: %v", err)
	}
	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		// TODO 404
		return u, fmt.Errorf("failed to findOne user by id, oid: %s, due to error: %v", id, err)
	}
	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decod user(id:%s) from DB due to error: %v", id, err)
	}

	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	panic("update")
}

func (d *db) Delete(ctx context.Context, id string) error {
	panic("delete")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
