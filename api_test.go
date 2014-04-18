package wpcom

import (
	"flag"
	"log"
	"testing"
	"time"

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
	return configTestClient(New())
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

func TestAnonNotes(t *testing.T) {
	c := getTestAnonymousClient()
	me, err := c.Me(false)
	notes, err := me.Notifications(O())
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if notes.LastSeen != 0 {
		t.Errorf("Expected notes.LastSeen of 0, got %d", notes.LastSeen)
	}
	if notes.Number != 0 {
		t.Errorf("Expected notes.Number of 0, got %d", notes.Number)
	}
}

func TestNotes(t *testing.T) {
	c := getTestClient()
	me, err := c.Me(false)
	notes, err := me.Notifications(O().Add("number", 3).Add("pretty", true))
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if notes.LastSeen == 0 {
		t.Errorf("Expected notes.LastSeen of !0, got %d", notes.LastSeen)
	}
	if notes.Number != 3 {
		t.Errorf("Expected notes.Number of 3, got %d", notes.Number)
	}
}

func TestNote(t *testing.T) {
	c := getTestClient()
	me, err := c.Me(false)
	note, err := me.Notification(1131732529)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if note.ID != 1131732529 {
		t.Errorf("Expected note.ID 1131732529, got %d", note.ID)
	}
}

func TestMissingNote(t *testing.T) {
	c := getTestClient()
	me, err := c.Me(false)
	note, err := me.Notification(-1)
	if err == nil {
		t.Errorf("Expected error, got %+v", note)
	}
}

func TestNoteSeen(t *testing.T) {
	c := getTestClient()
	me, err := c.Me(false)
	set, err := me.NotificationsSeen(time.Now().Unix())
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if set == false {
		t.Errorf("Expected true, got %#v", set)
	}
}

func TestNoteSeenBadTime(t *testing.T) {
	c := getTestClient()
	me, err := c.Me(false)
	set, err := me.NotificationsSeen(0)
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if set != false {
		t.Errorf("Expected false, got %#v", set)
	}
}

func TestNotesRead(t *testing.T) {
	c := getTestClient()
	me, err := c.Me(false)
	n, _ := me.Notifications(O().Add("number", 3))
	successes, err := me.NotificationsRead(map[int64]int64{
		n.Notifications[0].ID: -1,
		n.Notifications[1].ID: -1})
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	nsuccesses := 0
	for _, v := range successes {
		if v {
			nsuccesses++
		}
	}
	if nsuccesses != 2 {
		t.Errorf("Expected 2 successes, found %d (%+v)", nsuccesses, successes)
	}
}

func TestFreshlyPressed(t *testing.T) {
	c := getTestAnonymousClient()
	fp, err := c.FreshlyPressed()
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if fp.Posts[0].Title == "" {
		t.Errorf("Exptected a title, got ''")
	}
	if fp.Posts[0].URL == "" {
		t.Errorf("Exptected a URL, got ''")
	}
	if fp.Posts[0].Author.ID == 0 {
		t.Errorf("Expected an author ID, got 0")
	}
}

func TestSiteGetPosts(t *testing.T) {
	c := getTestClient()
	site, _ := c.SiteById(448698)
	resp, err := site.GetPosts(O().Add("number", 3))
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if len(resp.Posts) != 3 {
		t.Errorf("Expected 3 Found, got %d", len(resp.Posts))
	}
	if resp.Posts[0].Title == "" {
		t.Errorf("Expected a title for the first post returned. Got an empty string")
	}
	if resp.Posts[0].Author.Name == "" {
		t.Errorf("Expected an author name for the first post returned. Got an empty string")
	}
}

func TestSiteGetPostByID(t *testing.T) {
	c := getTestClient()
	site, _ := c.SiteById(448698)
	post, err := site.GetPost(113, O().Add("pretty", true))
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if post.ID != 113 {
		t.Errorf("Expected post.ID of 113, got %d", post.ID)
	}
	if post.Slug != "banzai" {
		t.Errorf("Expected post.Slug of 'banzai', got '%s'", post.Slug)
	}
}

func TestSiteGetPostBySlug(t *testing.T) {
	c := getTestClient()
	site, _ := c.SiteById(448698)
	post, err := site.GetPost("banzai", O().Add("pretty", true))
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if post.ID != 113 {
		t.Errorf("Expected post.ID of 113, got %d", post.ID)
	}
	if post.Slug != "banzai" {
		t.Errorf("Expected post.Slug of 'banzai', got '%s'", post.Slug)
	}
}

func TestSiteGetPostInvalidIDType(t *testing.T) {
	c := getTestClient()
	site, _ := c.SiteById(448698)
	post, err := site.GetPost(1.2345, O())
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
	if post != nil {
		t.Errorf("Expected nil, got %#v", post)
	}
}

func TestSiteGetPostInvalidID(t *testing.T) {
	c := getTestClient()
	site, _ := c.SiteById(448698)
	post, err := site.GetPost("nosuchpost", O())
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if post.ID != 0 {
		t.Errorf("Expected post.ID of 0, got %d", post.ID)
	}
	if post.Slug != "" {
		t.Errorf("Expected post.Slug of '', got '%s'", post.Slug)
	}
}
