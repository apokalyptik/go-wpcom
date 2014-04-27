package wpcom

import (
	"errors"
	"fmt"
)

// A site object to act upon
type Site struct {
	client       *Client
	ID           int                    `mapstructure:"ID"`
	Name         string                 `mapstructure:"name"`
	Description  string                 `mapstructure:"description"`
	URL          string                 `mapstructure:"URL"`
	Posts        int                    `mapstructure:"post_count"`
	Subscribers  int                    `mapstructure:"subscribers_count"`
	Lang         string                 `mapstructure:"lang"`
	Visible      string                 `mapstructure:"visible"`
	Options      map[string]SiteOptions `mapstructure:"options"`
	Meta         map[string]SiteMeta    `mapstructure:"meta"`
	Error        string                 `mapstructure:"error"`
	ErrorMessage string                 `mapstructure:"message"`
	Jetpack      bool                   `mapstructure:"jetpack"`
	Private      bool                   `mapstructure:"is_private"`
	Following    bool                   `mapstructure:"is_following"`
}

type SiteMeta map[string]string
type SiteOptions map[string]string

type SitePosts struct {
	Found int    `mapstructure:"found"`
	Posts []Post `mapstructure:"posts"`
}

// Get posts for a site.  For possible options see the following documentation
// URL:
// https://developer.wordpress.com/docs/api/1/get/sites/%24site/posts/
func (s *Site) GetPosts(o *Options) (rval *SitePosts, err error) {
	rval = new(SitePosts)
	prefix := fmt.Sprintf("sites/%d/posts/", s.ID)
	js, err := s.client.fetch(prefix, o, O())
	if err != nil {
		return
	}
	err = s.client.read(js, rval)
	for k, _ := range rval.Posts {
		rval.Posts[k].client = s.client.Clone()
	}
	return
}

// Get a site by ID, or by slug.  The function accepts both an integer type, or a
// string type for the id parameter.  Any other type will return an error and a nil
// reference.
// For possible options see the following documentation URLs:
// https://developer.wordpress.com/docs/api/1/get/sites/%24site/posts/%24post_ID/
// https://developer.wordpress.com/docs/api/1/get/sites/%24site/posts/slug:%24post_slug/
func (s *Site) GetPost(id interface{}, o *Options) (rval *Post, err error) {
	var prefix string
	rval = new(Post)
	switch v := id.(type) {
	case string:
		prefix = fmt.Sprintf("sites/%d/posts/slug:%s", s.ID, v)
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		prefix = fmt.Sprintf("sites/%d/posts/%d", s.ID, v)
	default:
		return nil, errors.New("Invalid id type")
	}
	js, err := s.client.fetch(prefix, o, O())
	if err != nil {
		return
	}
	err = s.client.read(js, rval)
	rval.client = s.client.Clone()
	return
}

// Query for comments on a site. See the following URL for possible options.
// http://developer.wordpress.com/docs/api/1/get/sites/%24site/comments/
func (s *Site) GetComments(o *Options) (comments *Comments, err error) {
	comments = new(Comments)
	prefix := fmt.Sprintf("sites/%d/comments/", s.ID)
	js, err := s.client.fetch(prefix, o, O())
	if err != nil {
		return
	}
	err = s.client.read(js, comments)
	return
}

// Query for a comment on a site. See the following URL for possible options.
// https://developer.wordpress.com/docs/api/1/get/sites/%24site/comments/%24comment_ID/
func (s *Site) Comment(id int) (comment *Comment, err error) {
	comment = new(Comment)
	prefix := fmt.Sprintf("sites/%d/comments/%d", s.ID, id)
	js, err := s.client.fetch(prefix, O(), O())
	if err != nil {
		return
	}
	err = s.client.read(js, comment)
	return
}
