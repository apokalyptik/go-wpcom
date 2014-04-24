package wpcom

import "testing"

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
