package db

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"os"
)

var client *firestore.Client

func getOption() option.ClientOption {
	rawString := os.Getenv("SERVICE_ACCOUNT_KEY")
	return option.WithCredentialsJSON([]byte(rawString))
}

func GetClient() (*firestore.Client, error) {
	if client == nil {

		ctx := context.Background()
		conf := &firebase.Config{ProjectID: "festa-notify"}

		app, err := firebase.NewApp(ctx, conf, getOption())

		if err != nil {
			return nil, fmt.Errorf("error initializing firebase app: %v", err)
		}

		store, err := app.Firestore(ctx)
		if err != nil {
			return nil, fmt.Errorf("error initializing firestore: %v", err)
		}

		client = store
	}
	return client, nil
}
