package wpcom

// A post object which is used for acting on a found post (ses the Site struct for finding posts)
type Post struct {
	client *Client
}

// Post details
func (p *Post) Info() {
}

// Modify the post
func (p *Post) Edit() {
}

// Delete the post
func (p *Post) Delete() {
}

// Get likes for the post
func (p *Post) GetLikes() {
}

// Like the post
func (p *Post) Like() {
}

// Unlike the post
func (p *Post) UnLike() {
}

// Reblog the post
func (p *Post) Reblog() {
}

// Whether the user has reblogged the post
func (p *Post) Reblogged() {
}

// Find related posts
func (p *Post) Related() {
}

// Get recent comments for the post
func (p *Post) Comments() {
}

// Comment on the post
func (p *Post) Comment() {
}
