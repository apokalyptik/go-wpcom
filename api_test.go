package wpcom

import (
	"flag"
	"log"

	"code.google.com/p/goconf/conf"
)

var testconf *conf.ConfigFile
var foundSitesForSiteByString = make(map[int]string)
var foundNotesForTesting []Notification

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
	return configTestClient(New())
}
