package wpcom

import "testing"

func TestSiteGetComments(t *testing.T) {
	c := getTestClient()
	site, _ := c.SiteById(448698)
	comments, err := site.GetComments(O())
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if len(comments.Comments) == 0 {
		t.Errorf("Expected at least one comment, got none")
	}
	if comments.Comments[0].ID == 0 {
		t.Errorf("Expected comment ID of !0, got 0")
	}
	if comments.Comments[0].Content == "" {
		t.Errorf("Expected non empty comment content, got ''")
	}
}
