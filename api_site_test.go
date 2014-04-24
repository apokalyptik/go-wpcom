package wpcom

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

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

func BenchmarkSiteById(b *testing.B) {
	c := getTestAnonymousClient()
	for i := 0; i < b.N; i++ {
		s, _ := c.SiteById(rand.Intn(100000) + 10000)
		if s.ID != 0 {
			foundSitesForSiteByString[len(foundSitesForSiteByString)] = s.URL
		}
	}
}

func BenchmarkSiteByName(b *testing.B) {
	c := getTestAnonymousClient()
	for i := 0; i < b.N; i++ {
		id := rand.Intn(len(foundSitesForSiteByString) - 1)
		siteUrl := foundSitesForSiteByString[id]
		siteHost := strings.TrimPrefix(strings.TrimPrefix(siteUrl, "http://"), "https://")
		c.SiteByString(siteHost)
	}
}

func BenchmarkSiteByBadName(b *testing.B) {
	c := getTestAnonymousClient()
	for i := 0; i < b.N; i++ {
		id := rand.Intn(len(foundSitesForSiteByString) - 1)
		siteUrl := foundSitesForSiteByString[id]
		siteHost := strings.TrimPrefix(strings.TrimPrefix(siteUrl, "http://"), "https://")
		c.SiteByString(fmt.Sprintf("%d.%s", rand.Int63(), siteHost))
	}
}
