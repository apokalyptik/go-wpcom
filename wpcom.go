// A client for accessing the WordPress.com (WPCOM) REST API V1
// See: http://developer.wordpress.com/docs/api/
package wpcom

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PREFIX = "https://public-api.wordpress.com/rest/v1/"

type client struct {
	token string
}

func (c *client) fetch(suffix string) (js []byte, err error) {
	url := fmt.Sprintf("%s%s", PREFIX, suffix)
	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	if c.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	resp, err := client.Do(req)
	if err != nil {
		return js, err
	}
	js, err = ioutil.ReadAll(resp.Body)
	return js, err
}

func softBool(input interface{}) (bool, error) {
	switch t := input.(type) {
	case int:
		if t < 1 {
			return false, nil
		}
		return true, nil
	case bool:
		return t, nil
	case string:
		if t == "" {
			return false, nil
		}
		return true, nil
	default:
		log.Printf("", input)
		return false, errors.New("Unhandled soft boolean type")

	}
}

func (c *client) read(js []byte, into interface{}) error {
	return json.Unmarshal(js, into)
}

// Generate a new WordPress.com REST API Client given an access token. See:
// https://developer.wordpress.com/docs/oauth2/
func New(access_token string) *client {
	client := new(client)
	client.token = access_token
	return client
}
