package wpcom

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Me struct {
	client       *Client
	ID           int                    `json:"ID"`
	DisplayName  string                 `json:"display_name"`
	Username     string                 `json:"username"`
	Email        string                 `json:"email"`
	BD           int                    `json:"email"`
	TokenSiteID  int                    `json:"token_site_id"`
	Avatar       string                 `json:"avatar_URL"`
	Profile      string                 `json:"profile_URL"`
	Verified     bool                   `json:"verified"`
	Meta         map[string]interface{} `json:"meta"`
	Error        string                 `json:"error"`
	ErrorMessage string                 `json:"message"`
}

func (m *Me) Get() error {
	js, err := m.client.fetch("me", O(), O())
	if err != nil {
		return err
	}
	err = m.client.read(js, &m)
	return err
}

func (m *Me) Notifications(opt *Options) (NotificationsResponse, error) {
	rval := NotificationsResponse{}
	js, err := m.client.fetch("notifications/", opt, O())
	if err != nil {
		return rval, err
	}
	err = m.client.read(js, &rval)
	rval.unhack()
	return rval, err
}

func (m *Me) Notification(id int64) (Notification, error) {
	rval := NotificationsResponse{}
	js, err := m.client.fetch(fmt.Sprintf("notifications/%d", id), O(), O())
	if err != nil {
		return Notification{}, err
	}
	err = m.client.read(js, &rval)
	rval.unhack()
	if rval.Number > 0 {
		return rval.Notifications[0], err
	} else {
		return Notification{}, errors.New("not found")
	}
}

func (m *Me) NotificationsSeen(timestamp int64) (success bool, err error) {
	js, err := m.client.fetch("notifications/seen", O(), O().Set("time", timestamp))
	if err != nil {
		return false, err
	}
	resp := make(map[string]interface{})
	err = m.client.read(js, &resp)
	if err != nil {
		return false, err
	}
	if v, ok := resp["success"]; ok {
		if v2, ok2 := resp["last_seen_time"]; ok2 {
			if hack(v2).int() == timestamp {
				return true, nil
			} else {
				return false, errors.New("response timestamp mismatched requested timestamp")
			}
		} else {
			return v.(bool), nil
		}
	} else {
		log.Printf("%s", resp)
	}
	return
}

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
	if hack(rval["success"]).bool() != true {
		return updated, errors.New("API returned failure result")
	}
	for _, v := range rval["updated"].([]interface{}) {
		k, _ := strconv.Atoi(v.(string))
		updated[int64(k)] = true
	}
	return
}
