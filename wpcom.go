// A client for accessing the WordPress.com (WPCOM) REST API V1
// See: http://developer.wordpress.com/docs/api/
package wpcom

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PREFIX = "https://public-api.wordpress.com/rest/v1/"

type Client struct {
	httpClient         *http.Client
	prefix             string
	token              string
	debug              bool
	insecureSkipVerify bool
}

func (c *Client) Me() (*Me, error) {
	rval := new(Me)
	rval.client = New(c.token)
	rval.client.Prefix(c.prefix)
	rval.client.Debug(c.debug)
	rval.client.InsecureSkipVerify(c.insecureSkipVerify)

	js, err := c.fetch("me", Options{})
	if err != nil {
		return rval, err
	}
	rval.raw = string(js)
	err = c.read(js, &rval)

	return rval, err
}

func (c *Client) SiteById(id int) (*Site, error) {
	return c.SiteByString(fmt.Sprintf("%d", id))
}

func (c *Client) SiteByString(hostname string) (*Site, error) {
	rval := new(Site)
	rval.client = New(c.token)
	rval.client.Prefix(c.prefix)
	rval.client.Debug(c.debug)
	rval.client.InsecureSkipVerify(c.insecureSkipVerify)

	js, err := c.fetch(fmt.Sprintf("sites/%s", hostname), Options{})
	if err != nil {
		return rval, err
	}
	rval.raw = string(js)
	err = c.read(js, &rval)
	return rval, err
}

// Set the URL Prefix for the API client. This should normally not change unless you are
// an Automattic developer with a WordPress.com development environment testing changes.
// This option should *never* be overridden outside this specific circumstance
func (c *Client) Prefix(prefix string) {
	c.prefix = prefix
}

// Enable or Disable Verification of the remote SSL Certificate. The client verifies by
// default, however, for Automattic developers with test environments the cert hostname
// does not match the request hostname. This function can be used to tell the client that
// This is OK. This option should *never* be disabled outside this specific circumstance
func (c *Client) InsecureSkipVerify(want bool) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: want},
	}
	c.httpClient = &http.Client{Transport: tr}
	c.insecureSkipVerify = want
}

// Turn debugging on or off.  When set to true request and response information will be
// logged using log.Printf()
func (c *Client) Debug(debug bool) {
	c.debug = debug
}

func (c *Client) fetch(suffix string, opt Options) (js []byte, err error) {
	var url string
	if getParams := opt.Encode(); getParams != "" {
		url = fmt.Sprintf("%s%s?%s", c.prefix, suffix, getParams)
	} else {
		url = fmt.Sprintf("%s%s", c.prefix, suffix)
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Host = "public-api.wordpress.com"
	if c.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	if c.debug {
		log.Printf("Request: %+v\n\nError: %+v", req, err)
	}
	resp, err := c.httpClient.Do(req)
	if c.debug {
		log.Printf("Response: %+v\n\nError: %+v", resp, err)
	}
	if err != nil {
		return js, err
	}
	js, err = ioutil.ReadAll(resp.Body)
	if c.debug {
		log.Printf("Response Text: %+v\n\nError: %+v", string(js), err)
	}
	return js, err
}

func (c *Client) read(js []byte, into interface{}) error {
	return json.Unmarshal(js, into)
}

// Generate a new WordPress.com REST API Client given an access token. See:
// https://developer.wordpress.com/docs/oauth2/
func New(access_token string) *Client {
	client := new(Client)
	client.prefix = PREFIX
	client.InsecureSkipVerify(false)
	client.token = access_token
	return client
}
