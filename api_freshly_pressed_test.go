package wpcom

import "testing"

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
