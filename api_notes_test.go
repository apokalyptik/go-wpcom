package wpcom

import (
	"math/rand"
	"testing"
	"time"
)

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
