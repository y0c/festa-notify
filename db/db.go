package db

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"os"
)

var client *db.Client

func getOption() option.ClientOption {
	rawString := os.Getenv("SERVICE_ACCOUNT_KEY")
	return option.WithCredentialsJSON([]byte(rawString))
}

func GetClient() (*db.Client, error) {
	if client == nil {
		ctx := context.Background()
		config := &firebase.Config{
			DatabaseURL: "https://festa-notify.firebaseio.com/",
		}

		app, err := firebase.NewApp(ctx, config, getOption())
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
