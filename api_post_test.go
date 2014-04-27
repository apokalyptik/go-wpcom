package wpcom

import "testing"

func TestPostComments(t *testing.T) {
	c := getTestClient()
	s, _ := c.SiteById(448698)
	p, _ := s.GetPost(40, O())
	comments, err := p.Comments(O())
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if comments.Found < 1 {
		t.Errorf("Expected one or more comments.Found, got 0")
	}
	if len(comments.Comments) < 1 {
		t.Errorf("Expected one or more comments, got 0")
	}
}
