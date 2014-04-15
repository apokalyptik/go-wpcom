package wpcom

// The structure of data representing a single notification. This is used
// for both Notifications() and Notification() calls on Me structs
type Notification struct {
	ID        int64                  `json:"-"`
	Type      string                 `json:"type"`
	Unread    int64                  `json:"-"`
	Noticon   string                 `json:"noticon"`
	Timestamp int64                  `json:"-"`
	Body      map[string]interface{} `json:"body"`
	Subject   map[string]interface{} `json:"subject"`
	// Hack* are an interrim solution for the things being documented
	// in the API as being returned as one type, but actually being
	// returned as another type.  These values are used for unmarshalling
	// into and then they're "fixed" with a call to an unexported "unhack"
	// method which puts, for example, an integer in ID from the String
	// actually stored in HackID.  These will come out if and when the API
	// is fixed. You should, therefor, never rely on this data and always
	// use the clean, typed, data.
	HackID        interface{} `json:"id"`
	HackUnread    interface{} `json:"unread"`
	HackTimestamp interface{} `json:"timestamp"`
}

func (n *Notification) unhack() {
	n.ID = hack(n.HackID).int()
	n.Timestamp = hack(n.HackTimestamp).int()
	n.Unread = hack(n.HackUnread).int()
}

// The struture for the data returned from a Notifications() call on a Me struct
type NotificationsResponse struct {
	Notifications []Notification `json:"notes"`
	LastSeen      int64          `json:"-"`
	Number        int            `json:"number"`
	// See the Documentation for Notification for information about Hack* members
	HackLastSeen interface{} `json:"last_seen_time"`
}

func (n *NotificationsResponse) unhack() {
	n.LastSeen = hack(n.HackLastSeen).int()
	for k, _ := range n.Notifications {
		n.Notifications[k].unhack()
	}
}
