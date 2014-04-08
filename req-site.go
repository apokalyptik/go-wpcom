package wpcom

import (
	"errors"
	"fmt"
)

type SiteResponse struct {
	ID            int                    `json:"ID"`
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	URL           string                 `json:"URL"`
	Posts         int                    `json:"post_count"`
	Subscribers   int                    `json:"subscribers_count"`
	Lang          string                 `json:"lang"`
	Visible       string                 `json:"visible"`
	Options       map[string]interface{} `json:"options"`
	Meta          map[string]interface{} `json:"meta"`
	Error         string                 `json:"error"`
	ErrorMessage  string                 `json:"message"`
	Raw           string                 `json:"-"`
	JetpackBool   interface{}            `json:"jetpack"`
	PrivateBool   interface{}            `json:"is_private"`
	FollowingBool interface{}            `json:"is_following"`
}

func (r *SiteResponse) Following() (bool, error) {
	if r.Error != "" {
		return false, errors.New(r.Error)
	}
	return softBool(r.FollowingBool)
}

func (r *SiteResponse) Jetpack() (bool, error) {
	if r.Error != "" {
		return false, errors.New(r.Error)
	}
	return softBool(r.JetpackBool)
}

func (r *SiteResponse) Private() (bool, error) {
	if r.Error != "" {
		return false, errors.New(r.Error)
	}
	return softBool(r.PrivateBool)
}

func (c *client) Site(site interface{}) (SiteResponse, error) {
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
	resp.Raw = string(js)
	err = c.read(js, &resp)
	return resp, err
}
