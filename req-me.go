package wpcom

type MeResponse struct {
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
	raw          string                 `json:"-"`
}

func (c *Client) Me() (MeResponse, error) {
	var resp MeResponse
	js, err := c.fetch("me")
	resp.raw = string(js)
	err = c.read(js, &resp)
	return resp, err
}
