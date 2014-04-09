package wpcom

import "fmt"

type SiteResponse struct {
	ID           int                    `json:"ID"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	URL          string                 `json:"URL"`
	Posts        int                    `json:"post_count"`
	Subscribers  int                    `json:"subscribers_count"`
	Lang         string                 `json:"lang"`
	Visible      string                 `json:"visible"`
	Options      map[string]interface{} `json:"options"`
	Meta         map[string]interface{} `json:"meta"`
	Error        string                 `json:"error"`
	ErrorMessage string                 `json:"message"`
	Jetpack      bool                   `json:"jetpack"`
	Private      bool                   `json:"is_private"`
	Following    bool                   `json:"is_following"`
	raw          string                 `json:"-"`
}

func (c *Client) Site(site interface{}) (SiteResponse, error) {
	var suffix string
	var resp SiteResponse
	switch t := site.(type) {
	case string:
		suffix = fmt.Sprintf("sites/%s", t)
	case int:
		suffix = fmt.Sprintf("sites/%d", t)
	}
	js, err := c.fetch(suffix)
	if err != nil {
		return resp, err
	}
	resp.raw = string(js)
	err = c.read(js, &resp)
	return resp, err
}
