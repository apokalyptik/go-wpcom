package wpcom

import (
	"flag"
	"log"
	"testing"

	"code.google.com/p/goconf/conf"
)

var testconf *conf.ConfigFile

func init() {
	var err error
	var cfongigFile string
	flag.StringVar(&cfongigFile, "cfg", "production.conf", "path to the config file for tests")
	flag.Parse()
	testconf, err = conf.ReadConfigFile(cfongigFile)
	if err != nil {
		log.Fatalf("Got error parsing test.conf: %s", err.Error())
	}
}

func configTestClient(c *Client) *Client {
	if testconf.HasOption("default", "prefix") {
		prefix, err := testconf.GetString("default", "prefix")
		if err != nil {
			log.Fatalf(err.Error())
		}
		if prefix != "" {
			c.Prefix(prefix)
			c.InsecureSkipVerify(true)
		}
	}
	return c
}

func getTestClient() *Client {
	key, err := testconf.GetString("user", "token")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return configTestClient(New(key))
}

func getTestAnonymousClient() *Client {
	return configTestClient(New(""))
}

func TestSiteByString(t *testing.T) {
	c := getTestAnonymousClient()
	r, e := c.Site("blog.apokalyptik.com")
	if e != nil {
		t.Errorf("Expected nil error, got: %s", e.Error())
	}
	if r.ID != 20645179 {
		t.Errorf("Got site ID %s, expected 20645179", r.ID)
	}
	if r.Jetpack != true {
		t.Errorf("Expected site to be a jetpack site")
	}
	if r.Private != false {
		t.Errorf("Expected site to be public")
	}
}

func TestSiteById(t *testing.T) {
	c := getTestAnonymousClient()
	r, e := c.Site(448698)
	if e != nil {
		t.Errorf("Expected nil error, got: %s", e.Error())
	}
	if r.ID != 448698 {
		t.Errorf("Expected")
	}
}

func TestBadSiteNotWp(t *testing.T) {
	c := getTestAnonymousClient()
	r, e := c.Site("amazon.com")
	if e != nil {
		t.Errorf("Expected nil error, got: %s", e.Error())
	}
	if r.Error != "unknown_blog" {
		t.Errorf("Expected Error: unknown_blog for amazon.com")
	}
}

func TestBadSiteInvalidId(t *testing.T) {
	c := getTestAnonymousClient()
	r, e := c.Site(-100)
	if e != nil {
		t.Errorf("Expexted nil error, got: %s", e.Error())
	}
	if r.Error != "unknown_blog" {
		t.Errorf("Expected Error: unknown_blog for site -100")
	}
}

func TestMe(t *testing.T) {
	c := getTestClient()
	r, e := c.Me()
	if e != nil {
		t.Errorf("Expected nil error, got %s", e.Error())
	}
	if r.Username != "apokalyptik" {
		t.Errorf("Expected Username 'apokalyptik', got '%s'", r.Username)
	}
	if r.Verified != true {
		t.Errorf("Expected verified user")
	}
}

func TestNotifications(t *testing.T) {
	// not done yet
	c := getTestClient()
	_, e := c.Notifications(Options{}.Add("number", "3").Add("ids", "1118447366").Add("ids", "1120655349"))
	if e != nil {
		t.Errorf("Expected nil error, got: %s", e.Error())
	}
}
