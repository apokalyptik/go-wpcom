package wpcom

type Post struct {
	client        *Client
	ID            int                       `mapstructure:"ID"`
	SiteId        int                       `mapstructure:"site_ID"`
	Author        PostAuthor                `mapstructure:"author"`
	Date          string                    `mapstructure:"date"`
	Modified      string                    `mapstructure:"modified"`
	Title         string                    `mapstructure:"title"`
	URL           string                    `mapstructure:"URL"`
	ShortURL      string                    `mapstructure:"short_URL"`
	Content       string                    `mapstructure:"content"`
	Excerpt       string                    `mapstructure:"excerpt"`
	Slug          string                    `mapstructure:"slug"`
	GUID          string                    `mapstructure:"guid"`
	Status        string                    `mapstructure:"status"`
	Password      string                    `mapstructure:"password"`
	Parent        bool                      `mapstructure:"parent"`
	CommentsOpen  bool                      `mapstructure:"comments_open"`
	LikeCount     int                       `mapstructure:"like_count"`
	ILike         bool                      `mapstructure:"i_like"`
	Reblogged     bool                      `mapstructure:"is_reblogged"`
	Following     bool                      `mapstructure:"is_following"`
	GlobalID      string                    `mapstructure:"global_ID"`
	FeaturedImage string                    `mapstructure:"featured_image"`
	Geo           bool                      `mapstructure:"mapstructure"`   // ?? maybe not bool?
	PublicizeURLs []string                  `mapstructure:"publicize_URLs"` // ?? maybe not strings?
	Tags          map[string]PostTag        `mapstructure:"tags"`
	Categories    map[string]PostCategories `mapstructure:"categories"`
	Attachments   map[int]PostAttachment    `mapstructure:"attachments"`
	Metadata      []PostMeta                `mapstructure:"metadata"`
	Meta          map[string]PostMeta       `mapstructure:"meta"`
	FeaturedMedia interface{}               `"mapstructure:"featured_media"`
}

type PostAttachment struct {
	ID       int    `mapstructure:"ID"`
	URL      string `mapstructure:"URL"`
	GUID     string `mapstructure:"guid"`
	MimeType string `mapstructure:"mime_type"`
	Width    int    `mapstructure:"width"`
	Height   int    `mapstructure:"height"`
}

type PostTag struct {
	ID          int                 `mapstructure:"ID"`
	Name        string              `mapstructure:"name"`
	Slug        string              `mapstructure:"slug"`
	Description string              `mapstructure:"description"`
	PostCount   int                 `mapstructure:"post_count"`
	Meta        map[string]PostMeta `mapstructure:"meta"`
}

type PostCategories struct {
	ID          int                 `mapstructure:"ID"`
	Name        string              `mapstructure:"name"`
	Slug        string              `mapstructure:"slug"`
	Description string              `mapstructure:"description"`
	PostCount   int                 `mapstructure:"post_count"`
	Parent      int                 `mapstructure:"parent"`
	Meta        map[string]PostMeta `mapstructure:"meta"`
}

type PostAuthor struct {
	ID         int
	Email      string `mapstructure:"email"`
	Name       string `mapstructure:"name"`
	NiceName   string `mapstructure:"nice_name"`
	URL        string `mapstructure:"URL"`
	AvatarURL  string `mapstructure:"avatar_URL"`
	ProfileURL string `mapstructure:"profile_url"`
	SiteID     int    `mapstructure:"site_ID"`
}

type PostMeta map[string]string
