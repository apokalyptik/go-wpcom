// A client for accessing the WordPress.com (WPCOM) REST API V1
// See: http://developer.wordpress.com/docs/api/
//
// Usage Example:
//  package main
//
//  import "github.com/apokalyptik/go-wpcom"
//  import "fmt"
//
//  func main() {
//      c := wpcom.New()
//      site, _ := c.SiteByString("en.blog.wordpress.com")
//      fmt.Printf("Site ID: %d\n", site.ID)
//  }
package wpcom

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const PREFIX = "https://public-api.wordpress.com/rest/v1/"

type Client struct {
	httpClient         *http.Client
	prefix             string
	token              string
	debug              bool
	insecureSkipVerify bool
}

// Create a new client duplicating the settings 1:1 for the current client
func (c *Client) Clone() *Client {
	rval := New(c.token)
	rval.Prefix(c.prefix)
	rval.Debug(c.debug)
	rval.InsecureSkipVerify(c.insecureSkipVerify)
	return rval
}

// Create a Me struct.  See the documentation for Me for more information
// about it and its members and methods. By default an API call is made to
// prepopulate information in the Me struct.  But for times when you don't
// actually need or want to make that call those would be wasted resources
// (cpu cycles, wall clock time, bandwidth, etc) and so you can disable this
// functionality by passing false to this method
func (c *Client) Me(fetch ...bool) (*Me, error) {
	rval := new(Me)
	rval.client = c.Clone()
	if len(fetch) > 0 && fetch[0] == false {
		return rval, nil
	}
	js, err := c.fetch("me", O(), O())
	if err != nil {
		return rval, err
	}
	err = c.read(js, &rval)

	return rval, err
}

// Fetch a site struct using the site's integer ID.
func (c *Client) SiteById(id int) (*Site, error) {
	return c.SiteByString(fmt.Sprintf("%d", id))
}

// Fetch a site struct using the site's string hostname.
func (c *Client) SiteByString(hostname string) (*Site, error) {
	rval := new(Site)
	rval.client = c.Clone()

	js, err := c.fetch(fmt.Sprintf("sites/%s", url.QueryEscape(hostname)), O(), O())
	if err != nil {
		return rval, err
	}
	err = c.read(js, &rval)
	return rval, err
}

// Fetch WordPress.com "Freshly Pressed" Posts.  See:
// https://developer.wordpress.com/docs/api/1/get/freshly-pressed/
func (c *Client) FreshlyPressed() (rval FreshlyPressedResponse, err error) {
	rval = FreshlyPressedResponse{}
	js, err := c.fetch("freshly-pressed", O().Add("pretty", true), O())
	err = c.read(js, &rval)
	return
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

func (c *Client) fetch(suffix string, getOptions *Options, postOptions *Options) (js []byte, err error) {
	var url string
	var req *http.Request
	var debug string
	if c.debug {
		debug = "API Request Debugging Information\n\n"
		getOptions.Set("pretty", true)
	}
	if false == getOptions.Empty() {
		url = fmt.Sprintf("%s%s?%s", c.prefix, suffix, getOptions.Encode())
	} else {
		url = fmt.Sprintf("%s%s", c.prefix, suffix)
	}
	if postOptions.Empty() {
		req, err = http.NewRequest("GET", url, nil)
	} else {
		req, err = http.NewRequest("POST", url, bytes.NewBufferString(postOptions.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(postOptions.Encode())))
	}
	req.Host = "public-api.wordpress.com"
	if c.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	resp, err := c.httpClient.Do(req)
	if c.debug {
		if false == postOptions.Empty() {
			debug = fmt.Sprintf(
				"%s---[ Req(err:%+v) ]---\n%+v\n---[ Post ]---\n%s\n---[ Resp ]---\n%+v",
				debug,
				err,
				req,
				postOptions.Encode(),
				resp)
		} else {
			debug = fmt.Sprintf(
				"%s---[ Req(err:%+v) ]---\n%+v\n---[ Resp ]---\n%+v",
				debug,
				err,
				req,
				resp)
		}
	}
	if err != nil {
		log.Printf(debug)
		return js, err
	}
	js, err = ioutil.ReadAll(resp.Body)
	if c.debug {
		debug = fmt.Sprintf("%s\n---[ js(err:%+v) ]---\n%s", debug, err, string(js))
		log.Printf(debug)
	}
	return js, err
}

func (c *Client) read(js []byte, into interface{}) error {
	return json.Unmarshal(js, into)
}

// Generate a new WordPress.com REST API Client given an access token. See:
// https://developer.wordpress.com/docs/oauth2/
func New(access_token ...string) *Client {
	client := new(Client)
	client.prefix = PREFIX
	client.InsecureSkipVerify(false)
	if len(access_token) > 0 {
		client.token = access_token[0]
	}
	return client
}
