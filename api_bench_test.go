package wpcom

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

import "math/rand"

var foundSitesForSiteByString = make(map[int]string)
var foundNotesForTesting []Notification

func BenchmarkAnonymousMe(b *testing.B) {
	c := getTestAnonymousClient()
	for i := 0; i < b.N; i++ {
		c.Me(true)
	}
}

func BenchmarkAuthedMe(b *testing.B) {
	c := getTestClient()
	for i := 0; i < b.N; i++ {
		c.Me(true)
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

func BenchmarkAnonNotes(b *testing.B) {
	c := getTestAnonymousClient()
	me, _ := c.Me(false)
	for i := 0; i < b.N; i++ {
		me.Notifications(O())
	}
}

func BenchmarkAuthedNotes(b *testing.B) {
	c := getTestClient()
	me, _ := c.Me(false)
	for i := 0; i < b.N; i++ {
		n, _ := me.Notifications(O())
		foundNotesForTesting = n.Notifications
	}
}

func BenchmarkNoteFetch(b *testing.B) {
	c := getTestClient()
	me, _ := c.Me(false)
	for i := 0; i < b.N; i++ {
		id := foundNotesForTesting[rand.Intn(len(foundNotesForTesting)-1)].ID
		me.Notification(id)
	}
}

func BenchmarkNotesSeen(b *testing.B) {
	c := getTestClient()
	me, _ := c.Me(false)
	for i := 0; i < b.N; i++ {
		me.NotificationsSeen(time.Now().Unix())
	}
}

func BenchmarkNotesSeenBad(b *testing.B) {
	c := getTestClient()
	me, _ := c.Me(false)
	for i := 0; i < b.N; i++ {
		me.NotificationsSeen(-1)
	}
}

func BenchmarkNoteRead(b *testing.B) {
	c := getTestClient()
	me, _ := c.Me(false)
	for i := 0; i < b.N; i++ {
		id := foundNotesForTesting[rand.Intn(len(foundNotesForTesting)-1)].ID
		me.NotificationsRead(map[int64]int64{id: -1})
	}
}
