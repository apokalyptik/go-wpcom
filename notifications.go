package wpcom

type Notification struct {
	HackID        interface{}            `json:"id"`
	ID            int64                  `json:"-"`
	Type          string                 `json:"type"`
	HackUnread    interface{}            `json:"unread"`
	Unread        int64                  `json:"-"`
	Noticon       string                 `json:"noticon"`
	HackTimestamp interface{}            `json:"timestamp"`
	Timestamp     int64                  `json:"-"`
	Body          map[string]interface{} `json:"body"`
	Subject       map[string]interface{} `json:"subject"`
}

func (n *Notification) unhack() {
	n.ID = hack(n.HackID).int()
	n.Timestamp = hack(n.HackTimestamp).int()
	n.Unread = hack(n.HackUnread).int()
}

type NotificationsResponse struct {
	Notifications []Notification `json:"notes"`
	HackLastSeen  interface{}    `json:"last_seen_time"`
	LastSeen      int64          `json:"-"`
	Number        int            `json:"number"`
}

func (n *NotificationsResponse) unhack() {
	n.LastSeen = hack(n.HackLastSeen).int()
	for k, _ := range n.Notifications {
		n.Notifications[k].unhack()
	}
}
