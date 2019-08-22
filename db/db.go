package db

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"fmt"
	"golang.org/x/net/context"
)

var client *db.Client

func GetClient() (*db.Client, error) {
	if client == nil {
		ctx := context.Background()
		config := &firebase.Config{
			DatabaseURL: "https://festa-notify.firebaseio.com/",
		}

		app, err := firebase.NewApp(ctx, config)
		if err != nil {
			return nil, fmt.Errorf("error initializing app: %v", err)
		}

		dbClient, err := app.Database(ctx)

		if err != nil {
			return nil, fmt.Errorf("error initializing database: %v", err)
		}
		client = dbClient
	}
	return client, nil
}
