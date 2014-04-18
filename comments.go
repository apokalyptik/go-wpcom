package wpcom

// The structure for a response to a query for multiple comments
type Comments struct {
	Found    int       `mapstructure:"found"`
	Comments []Comment `mapstructure:"comments"`
}

// The structure for a single comment
type Comment struct {
	ID       int                 `mapstructure:"ID"`
	Post     CommentPost         `mapstructure:"post"`
	Author   PostAuthor          `mapstructure:"author"`
	Date     string              `mapstructure:"date"`
	URL      string              `mapstructure:"URL"`
	ShortURL string              `mapstructure:"short_URL"`
	Content  string              `mapstructure:"content"`
	Status   string              `mapstructure:"status"`
	Parent   bool                `mapstructure:"parent"`
	Type     string              `mapstructure:"type"`
	Meta     map[string]PostMeta `mapstructure:"meta"`
}

// The structure for comment details about the post for which the comment was made
type CommentPost struct {
	ID   int    `mapstructure:"ID"`
	Type string `mapstructure:"type"`
	Link string `mapstructure:"link"`
}
