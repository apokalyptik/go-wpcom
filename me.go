package wpcom

import (
	"errors"
	"fmt"
	"strconv"
)

// A structure representing the user (or lack thereof for anonymous API usage)
// associated with the API calls.  Several API functions are attached to this
// structure when it makes sense for those API calls in no other context than
// that they are attached to an authenticated user. Notifications for example.
// You should not create a Me struct directly but instead use the Me() method
// on a Client struct. The reason for this is that a proper client for API
// calls will be embedded in the Me struct only when initialized in this fashion.
// The embedded client will retain the key debug, etc, settings from the creating
// Client.
type Me struct {
	client       *Client
	ID           int                    `mapstructure:"ID"`
	DisplayName  string                 `mapstructure:"display_name"`
	Username     string                 `mapstructure:"username"`
	Email        string                 `mapstructure:"email"`
	PrimaryBlog  int                    `mapstructure:"primary_blog"`
	TokenSiteID  int                    `mapstructure:"token_site_id"`
	Avatar       string                 `mapstructure:"avatar_URL"`
	Profile      string                 `mapstructure:"profile_URL"`
	Verified     bool                   `mapstructure:"verified"`
	Meta         map[string]interface{} `mapstructure:"meta"`
	Error        string                 `mapstructure:"error"`
	ErrorMessage string                 `mapstructure:"message"`
}

// Fetch, or re-fetch, details about the current API user.  The method updates
// the existing struct with new data where applicable from the results of this
// function call.  This is especially useful for Me structs initially created
// without fetching the user info (which is possible by passing false to the Me
// method on a Client struct)
// see: https://developer.wordpress.com/docs/api/1/get/me/
func (m *Me) Get() error {
	js, err := m.client.fetch("me", O(), O())
	if err != nil {
		return err
	}
	err = m.client.read(js, &m)
	return err
}

// Fetch the notifications for a user.  The opt argument allows you to
// specify query parameters to attach to the API call.  See
// https://developer.wordpress.com/docs/api/1/get/notifications/
// for possible options, and information about the data it returns
func (m *Me) Notifications(opt *Options) (NotificationsResponse, error) {
	rval := NotificationsResponse{}
	js, err := m.client.fetch("notifications/", opt, O())
	if err != nil {
		return rval, err
	}
	err = m.client.read(js, &rval)
	return rval, err
}

// Fetch information about a specific notification via it's note ID.
// See: https://developer.wordpress.com/docs/api/1/get/notifications/%24note_ID/
func (m *Me) Notification(id int64) (Notification, error) {
	rval := NotificationsResponse{}
	js, err := m.client.fetch(fmt.Sprintf("notifications/%d", id), O(), O())
	if err != nil {
		return Notification{}, err
	}
	err = m.client.read(js, &rval)
	if rval.Number > 0 {
		return rval.Notifications[0], err
	} else {
		return Notification{}, errors.New("not found")
	}
}

// Set the timestamp of the most recently seen notification for the current user
// See: https://developer.wordpress.com/docs/api/1/post/notifications/seen/
func (m *Me) NotificationsSeen(timestamp int64) (success bool, err error) {
	js, err := m.client.fetch("notifications/seen", O(), O().Set("time", timestamp))
	if err != nil {
		return false, err
	}
	resp := new(NotificationsSeenResponse)
	err = m.client.read(js, resp)
	if err != nil {
		return false, err
	}
	if resp.Success {
		if resp.LastSeen == timestamp {
			return true, nil
		}
		return false, errors.New("response timestamp mismatched requested timestamp")
	}
	return false, nil
}

// Mark a set of notifications as read.  The l map being passed matches the
// "counts" request parameter 1:1  from the Api documentation.  See:
// https://developer.wordpress.com/docs/api/1/post/notifications/read/
func (m *Me) NotificationsRead(l map[int64]int64) (updated map[int64]bool, err error) {
	updated = make(map[int64]bool)
	postOptions := new(Options)
	for k, v := range l {
		postOptions.Set(fmt.Sprintf("counts[%d]", k), v)
		updated[k] = false
	}
	js, err := m.client.fetch("notifications/read", O().Add("pretty", true), postOptions)
	if err != nil {
		return
	}
	rval := make(map[string]interface{})
	err = m.client.read(js, &rval)
	if rval["success"] != true {
		return updated, errors.New("API returned failure result")
	}
	for _, v := range rval["updated"].([]interface{}) {
		k, _ := strconv.Atoi(v.(string))
		updated[int64(k)] = true
	}
	return
}
