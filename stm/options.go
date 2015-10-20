package stm

func NewOptions() *Options {
	// Default values
	return &Options{
		"http://www.example.com",
		"", // http://s3.amazonaws.com/sitemap-generator/,
		"tmp/",
		"sitemaps/",
		// "sitemap",
		NewFileAdapter(),
	}
}

type Options struct {
	defaultHost  string
	sitemapsHost string
	publicPath   string
	sitemapsPath string
	// filename	 string
	adapter      Adapter
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

func (opts *Options) SetAdapter(adapter Adapter) {
	opts.adapter = adapter
}

func (opts *Options) SitemapsHost() string {
	if opts.sitemapsHost != "" {
		return opts.sitemapsHost
	}
	return opts.defaultHost
}
