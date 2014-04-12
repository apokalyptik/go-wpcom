package wpcom

type Me struct {
	client       *Client
	ID           int                    `json:"ID"`
	DisplayName  string                 `json:"display_name"`
	Username     string                 `json:"username"`
	Email        string                 `json:"email"`
	BlogID       int                    `json:"email"`
	TokenSiteID  int                    `json:"token_site_id"`
	Avatar       string                 `json:"avatar_URL"`
	Profile      string                 `json:"profile_URL"`
	Verified     bool                   `json:"verified"`
	Meta         map[string]interface{} `json:"meta"`
	Error        string                 `json:"error"`
	ErrorMessage string                 `json:"message"`
}

func (m *Me) Get() error {
	js, err := m.client.fetch("me", Options{})
	if err != nil {
		return err
	}
	err = m.client.read(js, &m)
	return err
}

func (m *Me) Notifications(opt Options) (NotificationsResponse, error) {
	rval := NotificationsResponse{}
	js, err := m.client.fetch("notifications/", opt)
	if err != nil {
		return rval, err
	}
	err = m.client.read(js, &rval)
	rval.unhack()
	return rval, err
}
