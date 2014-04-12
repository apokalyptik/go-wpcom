package wpcom

// A site object to act upon
type Site struct {
	client       *Client
	ID           int                    `json:"ID"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	URL          string                 `json:"URL"`
	Posts        int                    `json:"post_count"`
	Subscribers  int                    `json:"subscribers_count"`
	Lang         string                 `json:"lang"`
	Visible      string                 `json:"visible"`
	Options      map[string]interface{} `json:"options"`
	Meta         map[string]interface{} `json:"meta"`
	Error        string                 `json:"error"`
	ErrorMessage string                 `json:"message"`
	Jetpack      bool                   `json:"jetpack"`
	Private      bool                   `json:"is_private"`
	Following    bool                   `json:"is_following"`
}

// Get a site struct
func NewSite() {
}

// Get details about the site
func (s *Site) Info() {
}

// List recent posts
func (s *Site) GetPosts() {
}

// Get a specific post
func (s *Site) GetPost() {
}

// Make a new post
func (s *Site) Post() {}
