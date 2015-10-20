package stm

func NewOptions() *Options {
	// Default values
	return &Options{
		defaultHost:  "http://www.example.com",
		sitemapsHost: "", // http://s3.amazonaws.com/sitemap-generator/,
		publicPath:   "tmp/",
		sitemapsPath: "sitemaps/",
		filename:     "sitemap",
		verbose:      false,
		compress:     true,
	}
}

type Options struct {
	defaultHost  string
	sitemapsHost string
	publicPath   string
	sitemapsPath string
	filename     string
	verbose      bool
	compress     bool
}

func (opts *Options) SetDefaultHost(host string) {
	opts.defaultHost = host
}

func (opts *Options) SetSitemapsHost(host string) {
	opts.sitemapsPath = host
}

func (opts *Options) SetPublicPath(path string) {
	opts.publicPath = path
}

func (opts *Options) SetSitemapsPath(path string) {
	opts.sitemapsPath = path
}

func (opts *Options) SitemapsHost() string {
	if opts.sitemapsHost != "" {
		return opts.sitemapsHost
	}
	return opts.defaultHost
}
