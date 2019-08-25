package subscriber

import "testing"

func TestSubscriber(t *testing.T) {
	subscriberService, err := New()
	if err != nil {
		t.Errorf("error")
	}
	subscriberService.GetSubscribers()

}
