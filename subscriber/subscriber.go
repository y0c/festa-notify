package subscriber

import (
	fbdb "firebase.google.com/go/db"
	"fmt"
	"github.com/y0c/festa-notify/db"
	"golang.org/x/net/context"
)

type Subscriber struct {
	Mail     string
	Keywords []string
}

type Service struct {
	db *fbdb.Client
}

func New() (*Service, error) {
	client, err := db.GetClient()

	if err != nil {
		return nil, fmt.Errorf("clinet error %v", err)
	}

	return &Service{
		client,
	}, nil
}

func (s *Service) GetSubscribers() ([]Subscriber, error) {
	var subscribers []Subscriber
	query := s.db.NewRef("subscriber")

	if err := query.Get(context.Background(), &subscribers); err != nil {
		return nil, fmt.Errorf("error %v", err)
	}

	return subscribers, nil
}
