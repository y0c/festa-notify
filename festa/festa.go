package festa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Festa struct{}

type EventResponse struct {
	Page     string  `json:"page"`
	PageSize string  `json:"pageSize"`
	Total    int     `json:"total"`
	Rows     []Event `json:"rows"`
}

func New() *Festa {
	return &Festa{}
}

const API_ENDPOINT = "https://festa.io/api/v1/events"

func toQueryString(params map[string]string) string {
	arr := []string{}
	for k, v := range params {
		arr = append(arr, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(arr, "&")
}

func (f *Festa) GetEvents() (events []Event) {

	var eventResponse EventResponse
	queryParam := map[string]string{
		"page":                  "1",
		"pageSize":              "24",
		"order":                 "startDate",
		"excludeExternalEvents": "false",
	}
	resp, _ := http.Get(fmt.Sprintf("%s?%s", API_ENDPOINT, toQueryString(queryParam)))

	responseBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(responseBytes, &eventResponse)

	return eventResponse.Rows
}
