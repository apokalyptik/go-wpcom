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

func TestPostLikes(t *testing.T) {
	c := getTestClient()
	site, _ := c.SiteByString("en.blog.wordpress.com")
	post, _ := site.GetPost(28615, O())
	likes, err := post.Likes(O())
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if likes.Found < 1 {
		t.Errorf("Expected one or more likes.Found, got 0")
	}
	if len(likes.Likes) < 1 {
		t.Errorf("Expected one or more likes, got 0")
	}
}
