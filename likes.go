package wpcom

// The structure for a response to a query for multiple likes
type Likes struct {
	Found    int    `mapstructure:"found"`
	Comments []Like `mapstructure:"likes"`
	I_Like   bool   `mapstructure:"i_like"`
}

type Like struct {
	client         *Client
	ID             int    `mapstructure:"ID"`
	Author         string `mapstructure:"login"`
	Email          string `mapstructure:"email"`
	name           string `mapstructure:"modified"`
	nice_name      string `mapstructure:"title"`
	URL            string `mapstructure:"URL"`
	avatar_URL     string `mapstructure:"short_URL"`
	profile_URL    string `mapstructure:"content"`
	site_ID        string `mapstructure:"excerpt"`
	default_avatar string `mapstructure:"slug"`
}
