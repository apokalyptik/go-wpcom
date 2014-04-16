package wpcom

import (
	"time"
)

// A structure representing the response of a freshly-pressed API call response.
type FreshlyPressedResponse struct {
	DateRange FreshlyPressedDateRange `json:"date_range"`
	Number    int                     `json:"number"`
	Posts     []Object                `json:"posts"`
}

// A structure Representing the date range part of the responce from a
// freshly-pressed API call response
type FreshlyPressedDateRange struct {
	Newest time.Time `json:"newest"`
	Oldest time.Time `json:"oldest"`
}
