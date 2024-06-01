package main

// https://gohugo.io/content-management/front-matter/
type PostFrontmatter struct {
	Date        string     `yaml:"date"`
	Description string     `yaml:"description"`
	Title       string     `yaml:"title"`
	Params      PostParams `yaml:"params"`
}

type PostParams struct {
	Content    string   `yaml:"content"`
	FeedId     string   `yaml:"feed_id"`
	Id         string   `yaml:"id"`
	Link       string   `yaml:"link"`
	Categories []string `yaml:"categories"`
}

func NewPostFrontmatter(guid, link string) *PostFrontmatter {
	out := new(PostFrontmatter)
	out.Params.Id = buildSafeId(guid, link)
	out.Params.Link = link
	return out
}

func (f *PostFrontmatter) WithTitle(title string) {
	f.Title = truncateText(title, 200)
}

func (f *PostFrontmatter) WithCategories(cats []string) {
	f.Params.Categories = append(f.Params.Categories, cats...)
}

func (f *PostFrontmatter) WithDescription(description string) {
	f.Description = truncateText(readable(description), 200)
}

func (f *PostFrontmatter) WithDate(date string) {
	f.Date = date
}

func (f *PostFrontmatter) WithContent(content string) {
	f.Params.Content = truncateText(readable(content), 300)
}

func (f *PostFrontmatter) WithFeedLink(feed_link string) {
	f.Params.FeedId = buildSafeId("", feed_link)
}

type FeedFrontmatter struct {
	Date        string     `yaml:"date"`
	Description string     `yaml:"description"`
	Title       string     `yaml:"title"`
	Params      FeedParams `yaml:"params"`
}

type FeedParams struct {
	FeedLink   string   `yaml:"feedlink"`
	Id         string   `yaml:"id"`
	Link       string   `yaml:"link"`
	BlogRolls  []string `yaml:"blogrolls"`
	FeedType   string   `yaml:"feedtype"`
	Categories []string `yaml:"categories"`
}

func NewFeedFrontmatter(feed_url string) *FeedFrontmatter {
	out := new(FeedFrontmatter)
	out.Params.Id = buildSafeId("", feed_url)
	out.Params.FeedLink = feed_url
	return out
}

func (f *FeedFrontmatter) WithDate(date string) {
	f.Date = date
}

func (f *FeedFrontmatter) WithDescription(description string) {
	f.Description = truncateText(readable(description), 200)
}

func (f *FeedFrontmatter) WithTitle(title string) {
	f.Title = title
}

func (f *FeedFrontmatter) WithLink(link string) {
	f.Params.Link = link
}

func (f *FeedFrontmatter) WithFeedType(feedType string) {
	f.Params.FeedType = feedType
}

func (f *FeedFrontmatter) WithCategories(cats []string) {
	f.Params.Categories = append(f.Params.Categories, cats...)
}

func (f *FeedFrontmatter) WithBlogRolls(links []string) {
	f.Params.BlogRolls = links
}

type LinkFrontmatter struct {
	Params LinkParams `yaml:"params"`
}

type LinkParams struct {
	SourceType      NodeType `yaml:"source_type"`
	SourceURL       string   `yaml:"source_url"`
	DestinationType NodeType `yaml:"destination_type"`
	DestinationURL  string   `yaml:"destination_url"`
	LinkType        string   `yaml:"link_type"`
}

func NewLinkFrontmatter(source_type NodeType, source_url string, destination_type NodeType, destination_url, link_type string) *LinkFrontmatter {
	out := new(LinkFrontmatter)
	out.Params.SourceType = source_type
	out.Params.SourceURL = source_url
	out.Params.DestinationType = destination_type
	out.Params.DestinationURL = destination_url
	out.Params.LinkType = link_type
	return out
}

func buildLinkId(source, dest string) string {
	return md5Hex(source + "\n" + dest)
}
