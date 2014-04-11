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

func TestMe(t *testing.T) {
	c := getTestClient()
	me, err := c.Me()
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	id, err := testconf.GetInt("user", "userid")
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if me.ID != id {
		t.Errorf("Expected ID of %d, got %d", id, me.ID)
	}
}

func TestAnonMe(t *testing.T) {
	c := getTestAnonymousClient()
	me, err := c.Me()
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if me.ID != 0 {
		t.Errorf("Expected ID of 0, got %d", me.ID)
	}
}

func TestSiteString(t *testing.T) {
	c := getTestAnonymousClient()
	site, err := c.SiteByString("blog.apokalyptik.com")
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if site.ID != 20645179 {
		t.Errorf("Expected ID of 20645179, got %d", site.ID)
	}
}

func TestSiteBadString(t *testing.T) {
	c := getTestAnonymousClient()
	site, err := c.SiteByString("amazon.com")
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if site.ID != 0 {
		t.Errorf("Expected ID of 20645179, got %d", site.ID)
	}
}

func TestWpcomSiteId(t *testing.T) {
	c := getTestAnonymousClient()
	site, err := c.SiteById(448698)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if site.ID != 448698 {
		t.Errorf("Expected ID of 448698, got %d", site.ID)
	}
}

func TestWpcomSiteBadId(t *testing.T) {
	c := getTestAnonymousClient()
	site, err := c.SiteById(-1)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if site.ID != 0 {
		t.Errorf("Expected ID of 0, got %d", site.ID)
	}
}

/*
func TestNotifications(t *testing.T) {
	c := getTestClient()
	r, e := c.Notifications(Options{}.Add("number", "3"))
	if e != nil {
		t.Errorf("Expected nil error, got: %s", e.Error())
	}
	if r.Number != 3 {
		t.Errorf("Got %d notes, expected 3", r.Number)
	}
	r2, e2 := c.Notifications(
		Options{}.Add("ids[]", r.Notifications[1].ID).Add("ids[]", r.Notifications[2].ID))
	if e2 != nil {
		log.Printf("Expected no error, got: %s", e2.Error())
	}
	if r2.Number != 2 {
		t.Errorf("Got %d notes, expected 2", r2.Number)
	}
}
*/
