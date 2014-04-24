package wpcom

import "testing"

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
