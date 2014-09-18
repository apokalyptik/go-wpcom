package wpcom

// Likes is the structure for a response to a query for multiple post likes
type Likes struct {
	Found int    `mapstructure:"found"`
	Likes []Like `mapstructure:"likes"`
	ILike bool   `mapstructure:"i_like"`
}

// Like is the structure for a single post like
type Like struct {
	client        *Client
	ID            int    `mapstructure:"ID"`
	Author        string `mapstructure:"login"`
	Email         string `mapstructure:"email"`
	name          string `mapstructure:"modified"`
	niceName      string `mapstructure:"title"`
	URL           string `mapstructure:"URL"`
	avatarURL     string `mapstructure:"short_URL"`
	profileURL    string `mapstructure:"content"`
	siteID        string `mapstructure:"excerpt"`
	defaultAvatar string `mapstructure:"slug"`
}
