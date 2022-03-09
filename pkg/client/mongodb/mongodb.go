package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (*mongo.Database, error) {
	// mongoDBURL := "mongodb+srv://oybek:1102123112@cluster0.gs3vf.mongodb.net/userService?retryWrites=true&w=majority"
	var mongoDBURL string
	var isAuth bool

	if username == "" && password == "" {
		isAuth = false
		mongoDBURL = fmt.Sprintf("mongodb+srv://oybek:1102123112@cluster0.gs3vf.mongodb.net/userService?retryWrites=true&w=majority")
	} else {
		isAuth = true
		mongoDBURL = fmt.Sprintf("%s:%s@%s:%s", username, password, host, port)
	}

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}

	cleint, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("field to connect to MongoDB due to error: %v", err)
	}

	if err = cleint.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("field to ping to MongoDB due to error: %v", err)
	}

	return cleint.Database(database), nil
}
