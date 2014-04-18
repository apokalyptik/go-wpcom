package wpcom

import "fmt"

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
