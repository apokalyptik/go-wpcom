package wpcom

type Categories struct {
	Found      int        `mapstructure:"found"`
	Categories []Category `mapstructure:"categories"`
}

type Category struct {
	siteID      int
	client      *Client
	ID          int             `mapstructure:"ID"`
	Name        string          `mapstructure:"name"`
	Slug        string          `mapstructure:"slug"`
	Description string          `mapstructure:"description"`
	PostCount   int             `mapstructure:"post_count"`
	Parent      int             `mapstructure:"parent"`
	Meta        map[string]Meta `mapstructure:"meta"`
}
