package wpcom

// A structure representing the response of a freshly-pressed API call response.
type FreshlyPressedResponse struct {
	DateRange FreshlyPressedDateRange `mapstructure:"date_range"`
	Number    int                     `mapstructure:"number"`
	Posts     []Post                  `mapstructure:"posts"`
}

// A structure Representing the date range part of the responce from a
// freshly-pressed API call response
type FreshlyPressedDateRange struct {
	Newest string `mapstructure:"newest"`
	Oldest string `mapstructure:"oldest"`
}
