package subscriber

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/spf13/cast"
	"github.com/y0c/festa-notify/db"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"time"
)

type Subscriber struct {
	Ref           *firestore.DocumentRef
	Mail          string
	Keywords      []string
	LastCreatedAt time.Time
}

type Service struct {
	client *firestore.Client
	ctx    context.Context
}

const COLLECTION_ID = "subscribers"

func New() (*Service, error) {
	client, err := db.GetClient()
	ctx := context.Background()

	if err != nil {
		return nil, fmt.Errorf("clinet error %v", err)
	}

	return &Service{
		client,
		ctx,
	}, nil
}

func castToStringArray(data interface{}) (result []string) {
	interfaces := data.([]interface{})
	for _, v := range interfaces {
		result = append(result, v.(string))
	}
	return
}

func (s *Service) GetSubscribers() ([]Subscriber, error) {
	iter := s.client.Collection(COLLECTION_ID).Documents(s.ctx)
	var subscribers []Subscriber

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		subscriberMap := doc.Data()
		fmt.Println(subscriberMap)

		subscribers = append(subscribers,
			Subscriber{
				doc.Ref,
				cast.ToString(subscriberMap["mail"]),
				cast.ToStringSlice(subscriberMap["keyword"]),
				cast.ToTime(subscriberMap["lastCreatedAt"]),
			})
	}

	return subscribers, nil
}

func (s *Service) UpdateLastCreatedAt(ref *firestore.DocumentRef, createdAt time.Time) error {
	_, err := ref.Set(s.ctx, map[string]interface{}{
		"lastCreatedAt": createdAt,
	}, firestore.MergeAll)

	if err != nil {
		return fmt.Errorf("firestore update error %v", err)
	}
	return nil
}
