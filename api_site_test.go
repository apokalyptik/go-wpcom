package wpcom

import "testing"

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

func TestWpcomSiteID(t *testing.T) {
	c := getTestAnonymousClient()
	site, err := c.SiteById(448698)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if site.ID != 448698 {
		t.Errorf("Expected ID of 448698, got %d", site.ID)
	}
}

func TestWpcomSiteBadID(t *testing.T) {
	c := getTestAnonymousClient()
	site, err := c.SiteById(-1)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if site.ID != 0 {
		t.Errorf("Expected ID of 0, got %d", site.ID)
	}
}
