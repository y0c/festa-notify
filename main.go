package main

import (
	"github.com/joho/godotenv"
	"github.com/thoas/go-funk"
	"github.com/y0c/festa-notify/festa"
	"github.com/y0c/festa-notify/mail"
	"github.com/y0c/festa-notify/subscriber"
	"strings"
	"time"
)

func matchKeywordEvent(keywords []string) func(event festa.Event) bool {
	return func(event festa.Event) (isMatch bool) {
		for _, keyword := range keywords {
			if strings.Contains(event.Name, keyword) {
				return true
			}
		}
		return false
	}
}

func getSubscribers() []subscriber.Subscriber {
	subscriberService, err := subscriber.New()
	panicError(err)

	subscribers, err := subscriberService.GetSubscribers()
	panicError(err)
	return subscribers
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	_ = godotenv.Load()
	mail.Auth()

	subscribers := getSubscribers()
	festaAPI := festa.New()

	events := festaAPI.GetEvents()

	now := time.Now()

	availableEvents := funk.Filter(events, func(event festa.Event) bool {
		return now.Before(event.StartDate)
	}).([]festa.Event)

	for _, subscriber := range subscribers {
		personalEvents := funk.Filter(availableEvents, matchKeywordEvent(subscriber.Keywords)).([]festa.Event)
		m := mail.New([]string{subscriber.Mail}, "Festa 알림", personalEvents[0].Name)

		m.Send()
	}

}
