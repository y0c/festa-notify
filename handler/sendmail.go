package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/y0c/festa-notify/festa"
	"github.com/y0c/festa-notify/mail"
	"github.com/y0c/festa-notify/subscriber"
	"github.com/y0c/festa-notify/template"
	"net/http"
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

func SendMailHandler(c *gin.Context) {
	mail.Auth()

	subscribers := getSubscribers()
	festaAPI := festa.New()

	festaEvents := festaAPI.GetEvents()

	now := time.Now()

	availableEvents := funk.Filter(festaEvents, func(event festa.Event) bool {
		return now.Before(event.StartDate)
	}).([]festa.Event)

	fmt.Println(len(availableEvents))

	for _, subscriber := range subscribers {
		personalEvents := funk.Filter(availableEvents, matchKeywordEvent(subscriber.Keywords)).([]festa.Event)
		eventTemplate, err := template.GenerateEventTemplate(personalEvents)
		panicError(err)
		m := mail.New([]string{subscriber.Mail}, "Festa 알림", eventTemplate)

		_, err = m.Send()
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "OK"})
}
