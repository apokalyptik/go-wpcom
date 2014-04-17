package wpcom

import (
	"time"
)

// A structure representing the response of a freshly-pressed API call response.
type FreshlyPressedResponse struct {
	DateRange FreshlyPressedDateRange  `mapstructure:"date_range"`
	Number    int                      `mapstructure:"number"`
	Posts     []map[string]interface{} `mapstructure:"posts"`
}

// A structure Representing the date range part of the responce from a
// freshly-pressed API call response
type FreshlyPressedDateRange struct {
	Newest time.Time `mapstructure:"newest"`
	Oldest time.Time `mapstructure:"oldest"`
}
