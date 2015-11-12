package stm

func NewOptions() *Options {
	// Default values
	return &Options{
		defaultHost:  "http://www.example.com",
		sitemapsHost: "", // http://s3.amazonaws.com/sitemap-generator/,
		publicPath:   "public/",
		sitemapsPath: "sitemaps/",
		filename:     "sitemap",
		verbose:      true,
		compress:     true,
		adp:          NewFileAdapter(),
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
	adp          Adapter
	nmr          *Namer
	loc          *Location
}

func (opts *Options) SetDefaultHost(host string) {
	opts.defaultHost = host
}

func (opts *Options) SetSitemapsHost(host string) {
	opts.sitemapsHost = host
}

func (opts *Options) SetSitemapsPath(path string) {
	opts.sitemapsPath = path
}

func (opts *Options) SetPublicPath(path string) {
	opts.publicPath = path
}

func (opts *Options) SetFilename(filename string) {
	opts.filename = filename
}

func (opts *Options) SetVerbose(verbose bool) {
	opts.verbose = verbose
}

func (opts *Options) SetAdapter(adp Adapter) {
	opts.adp = adp
}

func (opts *Options) SitemapsHost() string {
	if opts.sitemapsHost != "" {
		return opts.sitemapsHost
	}
	return opts.defaultHost
}

func (opts *Options) Location() *Location {
	return NewLocation(opts)
}

func (opts *Options) IndexLocation() *Location {
	o := opts.Clone()
	o.nmr = NewNamer(&NOpts{base: opts.filename})
	return NewLocation(o)
}

func (opts *Options) Namer() *Namer {
	if opts.nmr == nil {
		opts.nmr = NewNamer(&NOpts{base: opts.filename, zero: 1, start: 2})
	}
	return opts.nmr
}

func (opts *Options) Clone() *Options {
	o := *opts
	return &o
}
