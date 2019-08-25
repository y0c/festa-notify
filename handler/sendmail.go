package handler

import (
	"github.com/thoas/go-funk"
	"github.com/y0c/festa-notify/festa"
	"github.com/y0c/festa-notify/mail"
	"github.com/y0c/festa-notify/subscriber"
	"github.com/y0c/festa-notify/template"
	"sort"
	"strings"
	"time"
)

func containsKeyword(event festa.Event, keyword string) bool {
	return strings.Contains(event.Name, keyword) || strings.Contains(event.HostOrganization.Name, keyword)
}

func matchKeywordEvent(keywords []string) func(event festa.Event) bool {
	return func(event festa.Event) (isMatch bool) {
		for _, keyword := range keywords {
			if containsKeyword(event, keyword) {
				return true
			}
		}
		return false
	}
}

func handleHTTPError(err error) {
	if err != nil {
		panic(err)
	}
}

// SendMailHandler lambda function handler for send notification to subscribers
func SendMailHandler() (string, error) {
	subscriberService, err := subscriber.New()
	handleHTTPError(err)

	subscribers, err := subscriberService.GetSubscribers()
	handleHTTPError(err)
	festaAPI := festa.New()
	festaEvents := festaAPI.GetEvents()
	now := time.Now()

	availableEvents := funk.Filter(festaEvents, func(event festa.Event) bool {
		return now.Before(event.StartDate)
	}).([]festa.Event)

	for _, subscriber := range subscribers {
		personalEvents := funk.Filter(availableEvents, matchKeywordEvent(subscriber.Keywords)).([]festa.Event)

		if !subscriber.LastCreatedAt.IsZero() {
			personalEvents = funk.Filter(personalEvents, func(event festa.Event) bool {
				return subscriber.LastCreatedAt.Before(event.CreatedAt)
			}).([]festa.Event)
		}

		if len(personalEvents) == 0 {
			continue
		}

		createdAts := funk.Map(personalEvents, func(event festa.Event) time.Time {
			return event.CreatedAt
		}).([]time.Time)

		sort.Slice(createdAts, func(i, j int) bool {
			return createdAts[i].After(createdAts[j])
		})

		subscriberService.UpdateLastCreatedAt(subscriber.Ref, createdAts[0])

		eventTemplate, err := template.GenerateEventTemplate(personalEvents)
		handleHTTPError(err)

		err = mail.Send(mail.Form{
			To:      subscriber.Mail,
			Body:    eventTemplate,
			Subject: "Festa 알림",
		})

		handleHTTPError(err)
	}
	return "Success", nil
}
