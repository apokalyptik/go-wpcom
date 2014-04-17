package wpcom

// The structure of data representing a single notification. This is used
// for both Notifications() and Notification() calls on Me structs
type Notification struct {
	ID        int64                  `mapstructure:"id"`
	Type      string                 `mapstructure:"type"`
	Unread    int64                  `mapstructure:"unread"`
	Noticon   string                 `mapstructure:"noticon"`
	Timestamp int64                  `mapstructure:"timestamp"`
	Body      map[string]interface{} `mapstructure:"body"`
	Subject   map[string]interface{} `mapstructure:"subject"`
}

// The struture for the data returned from a Notifications() call on a Me struct
type NotificationsResponse struct {
	Notifications []Notification `mapstructure:"notes"`
	LastSeen      int64          `mapstructure:"last_seen_time"`
	Number        int            `mapstructure:"number"`
}

type NotificationsSeenResponse struct {
	LastSeen int64 `mapstructure:"last_seen_time"`
	Success  bool  `mapstructure:"success"`
}
