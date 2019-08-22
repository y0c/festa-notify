package festa

import "time"

// Event ...
type Event struct {
	EventID             int         `json:"eventId"`
	HostUserID          interface{} `json:"hostUserId"`
	HostOrganizationID  int         `json:"hostOrganizationId"`
	Name                string      `json:"name"`
	EventSignature      string      `json:"eventSignature"`
	StartDate           time.Time   `json:"startDate"`
	EndDate             time.Time   `json:"endDate"`
	RefundDueDate       time.Time   `json:"refundDueDate"`
	Tag                 string      `json:"tag"`
	Description         string      `json:"description"`
	Charge              bool        `json:"charge"`
	Published           bool        `json:"published"`
	LimitWaiters        int         `json:"limitWaiters"`
	SurveyRequired      bool        `json:"surveyRequired"`
	SurveyBeforePayment bool        `json:"surveyBeforePayment"`
	External            bool        `json:"external"`
	ExternalLink        string      `json:"externalLink"`
	CreatedAt           time.Time   `json:"createdAt"`
	IsHostPicked        bool        `json:"isHostPicked"`
	Tickets             []struct {
		Registable    bool      `json:"registable"`
		TicketID      int       `json:"ticketId"`
		EventID       int       `json:"eventId"`
		Name          string    `json:"name"`
		Description   string    `json:"description"`
		Type          string    `json:"type"`
		Price         int       `json:"price"`
		Currency      string    `json:"currency"`
		Count         int       `json:"count"`
		Quantity      int       `json:"quantity"`
		LimitPerUser  int       `json:"limitPerUser"`
		SaleStartDate time.Time `json:"saleStartDate"`
		SaleEndDate   time.Time `json:"saleEndDate"`
		RefundDueDate time.Time `json:"refundDueDate"`
		HideRemains   bool      `json:"hideRemains"`
		UseSurvey     bool      `json:"useSurvey"`
		SurveyNotice  string    `json:"surveyNotice"`
	} `json:"tickets"`
	Location struct {
		LocationID  int    `json:"locationId"`
		EventID     int    `json:"eventId"`
		Name        string `json:"name"`
		Description string `json:"description"`
		CountryCode string `json:"countryCode"`
		State       string `json:"state"`
		City        string `json:"city"`
		PostalCode  string `json:"postalCode"`
		Address     string `json:"address"`
		Latitude    int    `json:"latitude"`
		Longitude   int    `json:"longitude"`
	} `json:"location"`
	Metadata struct {
		Contents    string `json:"contents"`
		BannerImage string `json:"bannerImage"`
		CoverImage  string `json:"coverImage"`
	} `json:"metadata"`
	HostOrganization struct {
		OrganizationID      int       `json:"organizationId"`
		Name                string    `json:"name"`
		Description         string    `json:"description"`
		ProfileImage        string    `json:"profileImage"`
		BannerImage         string    `json:"bannerImage"`
		HeaderImage         string    `json:"headerImage"`
		UseHeaderImage      bool      `json:"useHeaderImage"`
		DetailedDescription string    `json:"detailedDescription"`
		MainColor           string    `json:"mainColor"`
		Since               time.Time `json:"since"`
		CreatedAt           time.Time `json:"createdAt"`
	} `json:"hostOrganization"`
	HostUser interface{} `json:"hostUser"`
}
