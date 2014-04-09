package wpcom

type Notification struct {
	ID        uint64                 `json:"id"`
	Type      string                 `json:"type"`
	Unread    int                    `json:"unread"`
	Noticon   string                 `json:"noticon"`
	Timestamp int                    `json:"timestamp"`
	Body      map[string]interface{} `json:"body"`
	Subject   map[string]interface{} `json:"subject"`
	raw       string                 `json:"-"`
}

type NotificationsResponse struct {
	Notifications []Notification `json:"notes"`
	LastSeen      int            `json:"last_seen_time"`
	Number        int            `json:"number"`
}

func (c *Client) Notifications(opt Options) (NotificationsResponse, error) {
	rval := NotificationsResponse{}
	js, err := c.fetch("notifications/", opt)
	if err != nil {
		return rval, err
	}
	err = c.read(js, &rval)
	return rval, err
}
