package main

const USER_AGENT = "Feed2Pages/0.1"

const DEFAULT_READING_FOLDER = "reading"
const DEFAULT_FOLLOWING_FOLDER = "following"
const DEFAULT_DISCOVER_FOLDER = "discover"
const DEFAULT_NETWORK_FOLDER = "network"

const POST_PREFIX = "post-"
const FEED_PREFIX = "feed-"
const LINK_PREFIX = "link-"

// This uses a float as a workaround for Go-Colly
// marshaling, which converts int to float
type NodeType = float64

const (
	NODE_TYPE_SEED NodeType = iota
	NODE_TYPE_FEED
	NODE_TYPE_WEBSITE
	NODE_TYPE_BLOGROLL
)

type OutputMode = string

const (
	OUTPUT_MODE_HUGO_CONTENT OutputMode = "HugoContent"
	OUTPUT_MODE_SQL          OutputMode = "SQL"
)

var OUTPUT_MODES = []OutputMode{
	OUTPUT_MODE_HUGO_CONTENT,
	OUTPUT_MODE_SQL,
}

const (
	MIME_NONE     = ""
	MIME_TEXT_XML = "text/xml"
	MIME_APP_XML  = "application/xml"
	MIME_APP_ATOM = "application/atom+xml"
	MIME_APP_RSS  = "application/rss+xml"
	MIME_HTML     = "text/html"
	MIME_XHTML    = "application/xhtml+xml"
)

var OPML_MIMES = []string{
	MIME_TEXT_XML,
	MIME_APP_XML,
}

var FEED_MIMES = []string{
	MIME_APP_ATOM,
	MIME_APP_RSS,
}

var HTML_MIMES = []string{
	MIME_NONE,
	MIME_HTML,
	MIME_XHTML,
}

const (
	LINK_TYPE_LINK_REL_ME       = "rel_me"
	LINK_TYPE_LINK_REL_BLOGROLL = "rel_blogroll"
	LINK_TYPE_LINK_REL_ALT      = "rel_alt"
	LINK_TYPE_FROM_FEED         = "from_feed"
	LINK_TYPE_FROM_OPML         = "from_opml"
)

var META_ROBOT_NOINDEX_VARIANTS = []string{
	"noindex",
	"none",
}

var META_ROBOT_NOFOLLOW_VARIANTS = []string{
	"nofollow",
	"none",
}

const REFERER_STRING = "https://alexsci.com/rss-blogroll-network/"
