package wpcom

import (
	"time"
)

type FreshlyPressedResponse struct {
	DateRange FreshlyPressedDateRange `json:"date_range"`
	Number    int                     `json:"number"`
	Posts     []Object                `json:"posts"`
}

type FreshlyPressedDateRange struct {
	Newest time.Time `json:"newest"`
	Oldest time.Time `json:"oldest"`
}
